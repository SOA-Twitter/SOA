package handlers

import (
	"AuthService/data"
	"AuthService/proto/auth"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"regexp"
)

//var db = data.ConnectToDB()

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	l    *log.Logger
	repo data.AuthRepo
}

func NewAuthHandler(l *log.Logger, repo data.AuthRepo) *AuthHandler {
	return &AuthHandler{
		l:    l,
		repo: repo,
	}
}

func PasswordRegex(password string) error {
	pattern := regexp.MustCompile("[a-zA-Z]{8,}")

	res := pattern.MatchString(password)
	if res == true {
		return nil
	}
	return fmt.Errorf("Pattern error")
}

func (a *AuthHandler) Login(ctx context.Context, r *auth.LoginRequest) (*auth.LoginResponse, error) {
	a.l.Println("Login handler")
	res := &data.User{
		Username: r.Username,
		Password: r.Password,
	}
	err := a.repo.CheckCredentials(res.Username, res.Password)
	if err != nil {
		return &auth.LoginResponse{
			Status: http.StatusNotFound,
		}, err
	}
	userId, err := a.repo.FindUserID(res.Username)
	if err != nil {
		a.l.Println("Cannot find userId")
		return &auth.LoginResponse{
			Status: http.StatusNotFound,
		}, err
	}
	tokenString, _ := data.CreateJwt(userId, res.Username)
	return &auth.LoginResponse{
		Token:  tokenString,
		Status: http.StatusOK,
	}, nil
}

func (a *AuthHandler) Register(ctx context.Context, r *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	a.l.Println("Register handler")
	id := uuid.New().String()
	user := &data.User{
		UserId:   id,
		Username: r.Username,
		Password: r.Password,
	}

	result := PasswordRegex(user.Password)

	if result != nil {
		a.l.Println("Password must contain minimum eight characters, at least one uppercase letter," +
			" one lowercase letter, one number and one special character")
		return &auth.RegisterResponse{
			Status: http.StatusBadRequest,
		}, result
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		a.l.Println("Encryption failed", err)
		return &auth.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, result
	}
	user.Password = string(pass)

	err = a.repo.Register(user)
	if err != nil {
		a.l.Println("Error registration")
		return &auth.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, err
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
	return &auth.VerifyResponse{
		Status: http.StatusOK,
	}, nil
}
func (a *AuthHandler) GetUserId(ctx context.Context, r *auth.UserIdRequest) (*auth.UserIdResponse, error) {
	a.l.Println("Get UserID")
	claims := &data.Claims{}
	_, err := jwt.ParseWithClaims(r.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return data.SECRET, nil
	})
	if err != nil {
		return nil, err
	}
	return &auth.UserIdResponse{
		UserId: claims.UserId,
	}, nil
}
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello")))

}
