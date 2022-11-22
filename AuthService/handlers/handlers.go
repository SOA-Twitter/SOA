package handlers

import (
	"AuthService/data"
	"AuthService/proto/auth"
	"bufio"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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
		Email:    r.Email,
		Password: r.Password,
	}
	err := a.repo.CheckCredentials(res.Email, res.Password)
	if err != nil {
		return &auth.LoginResponse{
			Status: http.StatusNotFound,
		}, err
	}
	userId, err := a.repo.FindUserID(res.Email)
	if err != nil {
		a.l.Println("Cannot find userId")
		return &auth.LoginResponse{
			Status: http.StatusNotFound,
		}, err
	}
	tokenString, _ := data.CreateJwt(userId, res.Email)
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
		Email:    r.Email,
		Password: r.Password,
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

	err := a.repo.Register(user)
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
