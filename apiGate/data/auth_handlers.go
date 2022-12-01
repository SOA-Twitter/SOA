package data

import (
	"apiGate/protos/auth"
	"context"
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/mail"
	"regexp"
	"time"
)

type AuthHandler struct {
	l  *log.Logger
	pr auth.AuthServiceClient
}

type Gender string
type Role string

const (
	Male         Gender = "Male"
	Female       Gender = "Female"
	RegularUser  Role   = "RegularUser"
	BusinessUser Role   = "BusinessUser"
)

func (gender Gender) IsValid() error {
	switch gender {
	case Male, Female:
		return nil
	}
	return fmt.Errorf("invalid gender")
}

type User struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Gender         Gender `json:"gender"`
	Country        string `json:"country"`
	Age            int    `json:"age"`
	CompanyName    string `json:"company_name"`
	CompanyWebsite string `json:"company_website"`
	Role           Role   `json:"role"`
}
type ChangePass struct {
	OldPassword      string `json:"old_password"`
	NewPassword      string `json:"new_password"`
	RepeatedPassword string `json:"repeated_password"`
}
type Email struct {
	Email string `json:"email"`
}
type RecoverProfile struct {
	NewPassword      string `json:"new_password"`
	RepeatedPassword string `json:"repeated_password"`
}

func NewAuthHandler(l *log.Logger, pr auth.AuthServiceClient) *AuthHandler {
	return &AuthHandler{l, pr}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ah.l.Println("API-Gate - Login")

	var expirationTime = time.Now().Add(time.Second * 1200)
	user := User{}
	err := FromJSON(&user, r.Body)
	if err != nil {
		ah.l.Println("Cannot unmarshal json")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	token, err := ah.pr.Login(context.Background(), &auth.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		http.Error(w, "Invalid credentials!!!", http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token.Token,
		Path:    "/",
		Expires: expirationTime,
	})
	json.NewEncoder(w).Encode(token.Status)
}
func (ah *AuthHandler) SendRecoveryEmail(w http.ResponseWriter, r *http.Request) {
	ah.l.Println("API - Gate -Send Recovery Email")
	email := Email{}
	err := FromJSON(&email, r.Body)
	if err != nil {
		ah.l.Println("Cannot unmarshal json")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	res, err := ah.pr.SendRecoveryEmail(context.Background(), &auth.SendRecoveryEmailRequest{
		RecoveryEmail: email.Email,
	})
	if err != nil {
		ah.l.Println("Cannot recover profile")
		http.Error(w, "Cannot Recover profile. Error occurred", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res.Status)
}
func (ah *AuthHandler) RecoverProfile(w http.ResponseWriter, r *http.Request) {
	ah.l.Println("API-Gate - Recover profile")
	recProfil := RecoverProfile{}
	err := FromJSON(&recProfil, r.Body)
	if err != nil {
		ah.l.Println("Cannot unmarshal json")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	err3 := ValidatePassword(recProfil.NewPassword)
	if err3 != nil {
		ah.l.Println("Wrong password format")
		http.Error(w, "Password must be at least 8 characters long, with at least"+
			" one upper and one lower case letter, one special character and one number.", http.StatusBadRequest)
		return
	}
	if recProfil.NewPassword != recProfil.RepeatedPassword {
		ah.l.Println("Passwords do NOT match")
		http.Error(w, "Passwords do NOT match", http.StatusBadRequest)
		return
	}
	res, err := ah.pr.ResetPassword(context.Background(), &auth.ResetPasswordRequest{
		NewPassword:      recProfil.NewPassword,
		RepeatedPassword: recProfil.RepeatedPassword,
	})
	if err != nil {
		ah.l.Println("Cannot recover profile")
		http.Error(w, "Cannot Recover profile. Error occurred", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res.Status)
}

func (ah *AuthHandler) ActivateProfile(w http.ResponseWriter, r *http.Request) {
	ah.l.Println("API-Gate - Activate profile")
	activationUUID := mux.Vars(r)["activationId"]
	ah.l.Println("----------------------")
	ah.l.Println(activationUUID)
	ah.l.Println("----------------------")
	res, err := ah.pr.ActivateProfile(context.Background(), &auth.ActivationRequest{
		ActivationUUID: activationUUID,
	})
	if err != nil {
		ah.l.Println("Cannot activate profile")
		http.Error(w, "Cannot activate profile. Error occurred", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res.Status)
}

func (ah *AuthHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	ah.l.Println("API-Gate - Change Password")
	pass := ChangePass{}
	err := FromJSON(&pass, r.Body)
	if err != nil {
		ah.l.Println("Cannot unmarshal json")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	err3 := ValidatePassword(pass.NewPassword)
	if err3 != nil {
		ah.l.Println("Wrong password format")
		http.Error(w, "Password must be at least 8 characters long, with at least"+
			" one upper and one lower case letter, one special character and one number.", http.StatusBadRequest)
		return
	}
	if pass.NewPassword != pass.RepeatedPassword {
		ah.l.Println("Passwords do NOT match")
		http.Error(w, "Passwords do NOT match", http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized! NO COOKIE", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad request!", http.StatusBadRequest)
		return
	}
	res, err := ah.pr.ChangePassword(context.Background(), &auth.ChangePasswordRequest{
		OldPassword: pass.OldPassword,
		NewPassword: pass.NewPassword,
		Token:       c.Value,
	})
	if err != nil {
		json.NewEncoder(w).Encode(res.Status)
		http.Error(w, "Unable to save new password", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res.Status)

}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	ah.l.Println("API-Gate - Register")

	user := User{}
	err := FromJSON(&user, r.Body)
	if err != nil {
		ah.l.Println("Cannot unmarshal json")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	_, err1 := mail.ParseAddress(user.Email)
	_, err2 := regexp.MatchString("[a-zA-Z0-9._+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", user.Email)
	if err1 != nil || err2 != nil {
		ah.l.Println("Wrong email format")
		http.Error(w, "Wrong email format!.", http.StatusBadRequest)
		return
	}

	err3 := ValidatePassword(user.Password)
	if err3 != nil {
		ah.l.Println("Wrong password format")
		http.Error(w, "Password must be at least 8 characters long, with at least"+
			" one upper and one lower case letter, one special character and one number.", http.StatusBadRequest)
		return
	}

	if user.CompanyName != "" && user.CompanyWebsite != "" {

		_, error1 := regexp.MatchString("([a-zA-Z-']+)", user.CompanyName)
		if error1 != nil {
			ah.l.Println("Company name must not be empty")
			http.Error(w, "Company name must not be empty", http.StatusBadRequest)
			return
		}

		_, error2 := regexp.MatchString("([a-zA-Z-'_./]+)", user.CompanyWebsite)
		if error2 != nil {
			ah.l.Println("Company website must not be empty")
			http.Error(w, "Company website must not be empty", http.StatusBadRequest)
			return
		}

		res, err := ah.pr.Register(context.Background(), &auth.RegisterRequest{
			Email:          user.Email,
			Password:       user.Password,
			Username:       user.Username,
			CompanyName:    user.CompanyName,
			CompanyWebsite: user.CompanyWebsite,
			Role:           string(BusinessUser),
		})

		if err != nil {
			json.NewEncoder(w).Encode(res.Status)
			http.Error(w, "Registration unsuccessful", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(res.Status)
	} else {
		_, error1 := regexp.MatchString("([a-zA-Z-']{2,})", user.FirstName)
		if error1 != nil {
			ah.l.Println("First name must be at least 2 characters long")
			http.Error(w, "First name must be at least 2 characters long", http.StatusBadRequest)
			return
		}

		_, error2 := regexp.MatchString("([a-zA-Z-']{2,})", user.LastName)
		if error2 != nil {
			ah.l.Println("Last name must be at least 2 characters long")
			http.Error(w, "Last name must be at least 2 characters long", http.StatusBadRequest)
			return
		}

		_, error3 := regexp.MatchString("([a-zA-Z-]{4,})", user.Country)
		if error3 != nil {
			ah.l.Println("Country name must be at least 4 characters")
			http.Error(w, "Country name must be at least 4 characters", http.StatusBadRequest)
			return
		}

		_, error4 := regexp.MatchString("([a-zA-Z]{4,6})", string(user.Gender))

		if error4 != nil {
			ah.l.Println("Gender regex fail")
			http.Error(w, "Gender regex fail", http.StatusBadRequest)
			return
		} else {
			err4 := user.Gender.IsValid()
			if err4 != nil {
				ah.l.Println("Invalid gender")
				http.Error(w, "Invalid gender", http.StatusBadRequest)
				return
			}
		}

		if user.Age < 18 {
			ah.l.Println("User must have at least 18 years")
			http.Error(w, "User must have at least 18 years", http.StatusBadRequest)
			return
		}

		res, err := ah.pr.Register(context.Background(), &auth.RegisterRequest{
			Email:     user.Email,
			Password:  user.Password,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Gender:    string(user.Gender),
			Country:   user.Country,
			Age:       int32(user.Age),
			Role:      string(RegularUser),
		})

		if err != nil {
			json.NewEncoder(w).Encode(res.Status)
			http.Error(w, "Registration unsuccessful", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(res.Status)
	}
}

func HasPermission(userRole, path, method string) bool {
	e := casbin.NewEnforcer("data/rbac_model.conf", "data/policy.csv")
	if e.Enforce(userRole, path, method) {
		return true
	}
	return false
}

func (ah *AuthHandler) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ah.l.Println("Api-gate Middleware- Verify JWT")
		ah.l.Println(r.Method, r.URL.Path)
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized! NO COOKIE", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad request!", http.StatusBadRequest)
			return
		}
		tokenString := c.Value
		resp, err := ah.pr.VerifyJwt(context.Background(), &auth.VerifyRequest{
			Token: tokenString,
		})
		if err != nil || resp.Status != http.StatusOK {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}
		if !HasPermission(resp.UserRole, r.URL.Path, r.Method) {
			ah.l.Printf("User '%s' is not allowed to '%s' resource '%s'", resp.UserRole, r.Method, r.URL.Path)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)

	})
}
