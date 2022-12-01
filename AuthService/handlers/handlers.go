package handlers

import (
	"AuthService/data"
	"AuthService/proto/auth"
	"AuthService/proto/profile"
	"bufio"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

// *Standardization of INTENTION towards func SendEmail
const (
	INTENTION_RECOVERY   = "recovery"
	INTENTION_ACTIVATION = "activation"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	l    *log.Logger
	repo data.AuthRepo
	ps   profile.ProfileServiceClient
}

func NewAuthHandler(l *log.Logger, repo data.AuthRepo, ps profile.ProfileServiceClient) *AuthHandler {
	return &AuthHandler{
		l:    l,
		repo: repo,
		ps:   ps,
	}
}

func (a *AuthHandler) SendRecoveryEmail(ctx context.Context, r *auth.SendRecoveryEmailRequest) (*auth.RecoveryPasswordResponse, error) {
	a.l.Println("{authService} Send Recovery Email Handler")

	activationUUID, errEmailing := data.SendEmail(r.RecoveryEmail, INTENTION_RECOVERY)
	if errEmailing != nil {
		a.l.Println("RECOVERY Email delivery failed for: ", r.RecoveryEmail, errEmailing)
		return &auth.RecoveryPasswordResponse{
			Status: http.StatusBadRequest,
		}, errEmailing
	}

	errRecoveryReqSave := a.repo.SaveRecoveryRequest(activationUUID, r.RecoveryEmail)
	if errRecoveryReqSave != nil {
		a.l.Println("Error Writing Password Recovery Request to DB: ", errRecoveryReqSave)
		return &auth.RecoveryPasswordResponse{
			Status: http.StatusInternalServerError,
		}, errEmailing
	}

	return &auth.RecoveryPasswordResponse{
		Status: http.StatusOK,
	}, nil
}

func (a *AuthHandler) ResetPassword(ctx context.Context, r *auth.ResetPasswordRequest) (*auth.ChangePasswordResponse, error) {

	recoveryReq, errNotFound := a.repo.FindRecoveryRequest(r.RecoveryUUID)
	if errNotFound != nil {
		a.l.Println("Error Finding Password Recovery Request")
		return &auth.ChangePasswordResponse{
			Status: http.StatusNotFound,
		}, errNotFound
	}

	foundUser, errUserNotFound := a.repo.FindUser(recoveryReq.Email)
	if errUserNotFound != nil {
		a.l.Println("Error Finding User for email " + recoveryReq.Email)
		return &auth.ChangePasswordResponse{
			Status: http.StatusNotFound,
		}, errUserNotFound
	}

	// NEW PASSWORD GENERATING
	pass, err2 := bcrypt.GenerateFromPassword([]byte(r.NewPassword), bcrypt.DefaultCost)
	if err2 != nil {
		a.l.Println("Encryption failed for new PW", err2)
		return &auth.ChangePasswordResponse{
			Status: http.StatusInternalServerError,
		}, err2
	}
	foundUser.Password = string(pass)

	errPassChangeDatabase := a.repo.ChangePassword(foundUser.Email, foundUser.Password)
	if errPassChangeDatabase != nil {
		a.l.Println("Error Updating existing User (Password)")
		return &auth.ChangePasswordResponse{
			Status: http.StatusInternalServerError,
		}, errPassChangeDatabase
	}
	// NEW PASSWORD GENERATED

	errDelRecoveryReq := a.repo.DeleteRecoveryRequest(recoveryReq.RecoveryUUID, recoveryReq.Email)
	if errDelRecoveryReq != nil {
		a.l.Println("Error Deleting Password Recovery Request")
		//	*TODO: NEKAKAV ROLLBACK?
	}

	return &auth.ChangePasswordResponse{
		Status: http.StatusOK,
	}, nil

}

func (a *AuthHandler) ChangePassword(ctx context.Context, r *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	a.l.Println("{authService} - Change password handler")

	// find user by email from Cookie->Jwt CLAIMS; hash New-come Password & Compare to Old-db password; Write new Password hashed
	claims, err := data.GetFromClaims(r.Token)
	if err != nil {
		a.l.Println("Error getting claims")
		return &auth.ChangePasswordResponse{
			Status: http.StatusNotFound,
		}, err
	}
	foundUser, _ := a.repo.FindUser(claims.Email)

	err1 := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(r.OldPassword))

	if err1 != nil && err1 == bcrypt.ErrMismatchedHashAndPassword {
		a.l.Println("Invalid Password")
		return &auth.ChangePasswordResponse{
			Status: http.StatusBadRequest,
		}, err

	}

	pass, err2 := bcrypt.GenerateFromPassword([]byte(r.NewPassword), bcrypt.DefaultCost)
	if err2 != nil {
		a.l.Println("Encryption failed for new PW", err2)
		return &auth.ChangePasswordResponse{
			Status: http.StatusInternalServerError,
		}, err2
	}
	foundUser.Password = string(pass)

	err3 := a.repo.ChangePassword(foundUser.Email, foundUser.Password)
	if err3 != nil {
		a.l.Println("Error Updating existing User (Password)")
		return &auth.ChangePasswordResponse{
			Status: http.StatusInternalServerError,
		}, err3
	}
	return &auth.ChangePasswordResponse{Status: http.StatusOK}, nil

}
func (a *AuthHandler) ActivateProfile(ctx context.Context, r *auth.ActivationRequest) (*auth.ActivationResponse, error) {
	// Find {{KEY}} in DB, that equals to URL final section (activationUUID); Then set user.IsActivated = true, for user.Email == value of {{KEY}}
	// Finally delete the used acc. activation request from db
	a.l.Println("{authService} ActivateProfile Handler")

	activationReq, errNotFound := a.repo.FindActivationRequest(r.ActivationUUID)
	if errNotFound != nil {
		a.l.Println("Error Finding Account Activation Request")
		return &auth.ActivationResponse{
			Status: http.StatusNotFound,
		}, errNotFound
	}

	foundUser, errUserNotFound := a.repo.FindUser(activationReq.Email)
	if errUserNotFound != nil {
		a.l.Println("Error Finding User for email " + activationReq.Email)
		return &auth.ActivationResponse{
			Status: http.StatusNotFound,
		}, errUserNotFound
	}

	errActivating := a.repo.Edit(foundUser.Email)
	if errActivating != nil {
		a.l.Println("Error Editing User (Activating Account)")
		return &auth.ActivationResponse{
			Status: http.StatusInternalServerError,
		}, errActivating
	}
	errDelAccActReq := a.repo.DeleteActivationRequest(activationReq.ActivationUUID, activationReq.Email)
	if errDelAccActReq != nil {
		a.l.Println("Error Deleting Account Activation Request")
		//	*TODO: NEKAKAV ROLLBACK?
	}

	return &auth.ActivationResponse{
		Status: http.StatusOK,
	}, nil
}

func (a *AuthHandler) Login(ctx context.Context, r *auth.LoginRequest) (*auth.LoginResponse, error) {
	a.l.Println("Login handler")
	res := &data.User{
		Email:    r.Email,
		Password: r.Password,
	}
	err := a.repo.CheckCredentials(res.Email, res.Password)
	if err != nil {
		return &auth.LoginResponse{
			Status: http.StatusNotFound,
		}, err
	}
	email, role, err2 := a.repo.FindUserEmail(res.Email)
	if err2 != nil {
		a.l.Println("Cannot find user")
		return &auth.LoginResponse{
			Status: http.StatusNotFound,
		}, err2
	}
	a.l.Println(role)
	tokenString, _ := data.CreateJwt(email, role)
	return &auth.LoginResponse{
		Token:  tokenString,
		Status: http.StatusOK,
	}, nil

}

func (a *AuthHandler) Register(ctx context.Context, r *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	a.l.Println("Register handler")
	user := &data.User{
		Email:       r.Email,
		Password:    r.Password,
		Role:        r.Role,
		IsActivated: false,
	}
	_, _, err := a.repo.FindUserEmail(user.Email)
	if err == nil {
		a.l.Println("Email already exists")
		return &auth.RegisterResponse{
			Status: http.StatusBadRequest,
		}, err
	}
	file, err1 := os.Open("10k-most-common.txt")
	if err1 != nil {
		a.l.Println("Error opening 10k-most-common.txt file")
		a.l.Println(err1)
	}
	defer file.Close()

	var passwords []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	for _, v := range passwords {
		if user.Password == v {
			return &auth.RegisterResponse{
				Status: http.StatusBadRequest,
			}, nil
		}
	}

	pass, err2 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err2 != nil {
		a.l.Println("Encryption failed", err2)
		return &auth.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, err2
	}
	user.Password = string(pass)

	err3 := a.repo.Register(user)
	if err3 != nil {
		a.l.Println("Error registration")
		return &auth.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, err3
	}
	a.l.Println("______________________________")
	a.l.Println(user.Role)

	activationUUID, errEmailing := data.SendEmail(user.Email, INTENTION_ACTIVATION)
	if errEmailing != nil {
		a.l.Println("ACTIVATION Email delivery failed for: ", r.Email, errEmailing)
		error1 := a.repo.Delete(r.Email)
		if error1 != nil {
			a.l.Println("Error deleting user with email " + user.Email)
			return &auth.RegisterResponse{
				Status: http.StatusInternalServerError,
			}, error1
		}
		a.l.Println("Email delivery failed for: ", user.Email, errEmailing)
		return &auth.RegisterResponse{
			Status: http.StatusBadRequest,
		}, errEmailing
	}

	errActivationReqSave := a.repo.SaveActivationRequest(activationUUID, user.Email)
	if errActivationReqSave != nil {
		error1 := a.repo.Delete(r.Email)
		if error1 != nil {
			a.l.Println("Error deleting user with email " + user.Email)
			return &auth.RegisterResponse{
				Status: http.StatusInternalServerError,
			}, error1
		}
		a.l.Println("Error Writing Account Activation Request to DB: ", errActivationReqSave)
		return &auth.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, errActivationReqSave
	}

	_, err4 := a.ps.Register(context.Background(), &profile.ProfileRegisterRequest{
		Email:          r.Email,
		Username:       r.Username,
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		Gender:         r.Gender,
		Country:        r.Country,
		Age:            r.Age,
		CompanyWebsite: r.CompanyWebsite,
		CompanyName:    r.CompanyName,
	})

	if err4 != nil {
		// Also Deleting the account Activation combination (activationuuid + email) from DB, since profileSvc failed

		errDelAccActReq := a.repo.DeleteActivationRequest(activationUUID, user.Email)
		if errDelAccActReq != nil {
			a.l.Println("Error deleting Acc. Activation Request for {activationUUID, email} " + activationUUID + ", " + user.Email)
			return &auth.RegisterResponse{
				Status: http.StatusInternalServerError,
			}, errDelAccActReq
		}

		error1 := a.repo.Delete(r.Email)
		if error1 != nil {
			a.l.Println("Error deleting user with email " + user.Email)
			return &auth.RegisterResponse{
				Status: http.StatusInternalServerError,
			}, error1
		}
		a.l.Println("Deleting user with email " + user.Email)
		return &auth.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, err4
	}

	return &auth.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (a *AuthHandler) VerifyJwt(ctx context.Context, r *auth.VerifyRequest) (*auth.VerifyResponse, error) {
	a.l.Println("Verify JWT")
	err := data.ValidateJwt(r.Token)
	if err != nil {
		a.l.Println("JWT expired")
		return &auth.VerifyResponse{
			Status: http.StatusUnauthorized,
		}, nil
	}
	claims := &data.Claims{}
	_, err = jwt.ParseWithClaims(r.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return data.SampleSecretKey, nil
	})
	return &auth.VerifyResponse{
		Status:   http.StatusOK,
		UserRole: claims.Role,
	}, nil
}

func (a *AuthHandler) GetUser(ctx context.Context, r *auth.UserRequest) (*auth.UserResponse, error) {
	a.l.Println("Get User from JWT")
	claims, err := data.GetFromClaims(r.Token)
	if err != nil {
		return nil, err
	}
	return &auth.UserResponse{
		UserEmail: claims.Email,
		UserRole:  claims.Role,
	}, nil
}
