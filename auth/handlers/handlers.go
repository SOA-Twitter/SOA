package handlers

import (
	"TweeterMicro/auth/data"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

//var db = data.ConnectToDB()

type AuthHandler struct {
	l    *log.Logger
	repo data.AuthRepo
}

func NewAuthHandler(l *log.Logger, repo data.AuthRepo) *AuthHandler {
	return &AuthHandler{l, repo}
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	a.l.Println("Login handler")
	expirationTime := time.Now().Add(time.Second * 1200)
	user := &data.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, "Bad JSON format", http.StatusBadRequest)
		return
	}
	err = a.repo.FindUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Wrong credentials!", http.StatusBadRequest)
		return
	}
	tokenString, _ := data.CreateJwt(user.Username)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	json.NewEncoder(w).Encode(tokenString)
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	a.l.Println("Register handler")
	user := &data.User{}
	json.NewDecoder(r.Body).Decode(user)
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		a.l.Println("Encryption failed", err)
		return
	}
	user.Password = string(pass)

	err = a.repo.Register(user)
	if err != nil {
		a.l.Println("Error registration")
	}

}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello")))

}
