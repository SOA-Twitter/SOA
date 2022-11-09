package handlers

import (
	"TweeterMicro/AuthService/data"
	"TweeterMicro/AuthService/proto/auth"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
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

func (a *AuthHandler) Login(ctx context.Context, r *auth.LoginRequest) (*auth.LoginResponse, error) {
	a.l.Println("Login handler")
	res := &data.User{
		Username: r.Username,
		Password: r.Password,
	}
	err := a.repo.FindUser(res.Username, res.Password)
	if err != nil {
		return &auth.LoginResponse{
			Status: http.StatusNotFound,
		}, err
	}
	tokenString, _ := data.CreateJwt(res.Username)
	return &auth.LoginResponse{
		Token:  tokenString,
		Status: http.StatusOK,
	}, nil
}

func (a *AuthHandler) Register(ctx context.Context, r *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	a.l.Println("Register handler")
	user := &data.User{
		Username: r.Username,
		Password: r.Password,
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		a.l.Println("Encryption failed", err)
		return &auth.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, err
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
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello")))

}
