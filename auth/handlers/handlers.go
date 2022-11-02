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

var db = data.ConnectToDB()

type AuthHandler struct {
	l *log.Logger
}

func NewAuthHandler(l *log.Logger) *AuthHandler {
	return &AuthHandler{l}
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
	tokenString, err := FindOne(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Invalid credentials!", http.StatusUnauthorized)
		a.l.Println("Invalid credentials!")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	json.NewEncoder(w).Encode(tokenString)
}
func FindOne(username, password string) (string, error) {
	user := &data.User{}
	log.Println(user.Username)
	if err := db.Where("Username = ?", username).First(user).Error; err != nil {
		log.Println("Invalid Email")
		return "", err
	}
	log.Println(user.Username)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println("Invalid Password")
		return "", err

	}
	token, _ := data.CreateJwt(username)

	return token, nil
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
	createdUser := db.Create(user)
	var errMessage = createdUser.Error
	if createdUser.Error != nil {
		fmt.Println(errMessage)
		a.l.Println("Unable to Create user", errMessage)

	}
	json.NewEncoder(w).Encode(createdUser)
	w.WriteHeader(http.StatusCreated)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello")))

}
