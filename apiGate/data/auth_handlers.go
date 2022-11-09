package data

import (
	"TweeterMicro/AuthService/proto/auth"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
	l  *log.Logger
	pr auth.AuthServiceClient
}
type UserDAO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthHandler(l *log.Logger, pr auth.AuthServiceClient) *AuthHandler {
	return &AuthHandler{l, pr}
}
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var expirationTime = time.Now().Add(time.Second * 1200)
	user := UserDAO{}
	err := FromJSON(&user, r.Body)
	if err != nil {
		ah.l.Println("Cannot unmarshal json")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	token, err := ah.pr.Login(context.Background(), &auth.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token.Token,
		Expires: expirationTime,
	})
	json.NewEncoder(w).Encode(token.Status)
}
func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := UserDAO{}
	err := FromJSON(&user, r.Body)
	if err != nil {
		ah.l.Println("Cannot unmarshal json")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	res, err := ah.pr.Register(context.Background(), &auth.RegisterRequest{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		json.NewEncoder(w).Encode(res.Status)
		return
	}
	json.NewEncoder(w).Encode(res.Status)

}
