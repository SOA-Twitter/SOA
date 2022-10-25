package handlers

import (
	"TweeterMicro/auth/data"
	"TweeterMicro/auth/model"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
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
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	a.l.Println("Register handler")
	user := &model.User{}
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

func FindOne(email, password string) map[string]interface{} {
	user := &model.User{}
	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp

	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credential"}
		return resp
	}
	token, _ := data.CreateJwt(email, password)
	var response = map[string]interface{}{"status": false, "message": "logged in"}
	response["token"] = token
	return response
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello")))

}
