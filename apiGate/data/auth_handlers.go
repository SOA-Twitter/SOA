package data

import (
	"apiGate/protos/auth"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"time"
)

type AuthHandler struct {
	l  *log.Logger
	pr auth.AuthServiceClient
}
type UserDAO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(l *log.Logger, pr auth.AuthServiceClient) *AuthHandler {
	return &AuthHandler{l, pr}
}
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ah.l.Println("API-Gate - Login")

	var expirationTime = time.Now().Add(time.Second * 1200)
	user := UserDAO{}
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
func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	ah.l.Println("API-Gate - Register")

	user := UserDAO{}
	err := FromJSON(&user, r.Body)
	if err != nil {
		ah.l.Println("Cannot unmarshal json")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		ah.l.Println("Wrong email format")
		http.Error(w, "Wrong email format!.", http.StatusBadRequest)
		return
	}
	result := ValidatePassword(user.Password)

	if result != nil {
		ah.l.Println("Wrong password format")
		http.Error(w, "Password must be at least 7 characters long, with at least"+
			" one upper and one lower case letter, one special character and one number.", http.StatusBadRequest)
		return
	}
	res, err := ah.pr.Register(context.Background(), &auth.RegisterRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		json.NewEncoder(w).Encode(res.Status)
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res.Status)

}
func (ah *AuthHandler) VerifyJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ah.l.Println("Api-gate Middleware- Verify JWT")
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
		status, err := ah.pr.VerifyJwt(context.Background(), &auth.VerifyRequest{
			Token: tokenString,
		})
		if err != nil || status.Status != http.StatusOK {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)

	})
}
