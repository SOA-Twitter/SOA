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

//var db = data.ConnectToDB()

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
		Email:    r.Email,
		Password: r.Password,
		Role:     r.Role,
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
	claims := &data.Claims{}
	_, err := jwt.ParseWithClaims(r.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return data.SampleSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return &auth.UserResponse{
		UserEmail: claims.Email,
		UserRole:  claims.Role,
	}, nil
}
