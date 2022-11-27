package data

import (
	"apiGate/protos/auth"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"time"
)

type AuthHandler struct {
	l  *log.Logger
	pr auth.AuthServiceClient
}

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

func (gender Gender) IsValid() error {
	switch gender {
	case Male, Female:
		return nil
	}
	return fmt.Errorf("invalid gender")
}

type User struct {
	Username  string `json:"username" `
	Email     string `json:"email" `
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    Gender `json:"gender"`
	Country   string `json:"country"`
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
	if err1 != nil {
		ah.l.Println("Wrong email format")
		http.Error(w, "Wrong email format!.", http.StatusBadRequest)
		return
	}

	err3 := ValidatePassword(user.Password)
	if err3 != nil {
		ah.l.Println("Wrong password format")
		http.Error(w, "Password must be at least 7 characters long, with at least"+
			" one upper and one lower case letter, one special character and one number.", http.StatusBadRequest)
		return
	}

	if user.FirstName == "" {
		ah.l.Println("First name cannot be empty")
		http.Error(w, "First name cannot be empty", http.StatusBadRequest)
		return
	}

	if user.LastName == "" {
		ah.l.Println("Last name cannot be empty")
		http.Error(w, "Last name cannot be empty", http.StatusBadRequest)
		return
	}

	if user.Country == "" {
		ah.l.Println("Country cannot be empty")
		http.Error(w, "Country cannot be empty", http.StatusBadRequest)
		return
	}

	if user.Gender == "" {
		ah.l.Println("Gender cannot be empty")
		http.Error(w, "Gender cannot be empty", http.StatusBadRequest)
		return
	}

	err4 := user.Gender.IsValid()
	if err4 != nil {
		ah.l.Println("Invalid gender")
		http.Error(w, "Invalid gender", http.StatusBadRequest)
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
	})

	if err != nil {
		json.NewEncoder(w).Encode(res.Status)
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res.Status)

}

func (ah *AuthHandler) Authorize(next http.Handler) http.Handler {
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
		resp, err := ah.pr.VerifyJwt(context.Background(), &auth.VerifyRequest{
			Token: tokenString,
		})
		if err != nil || resp.Status != http.StatusOK {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}

		//OVDE DOBIJEMO ROLU
		//i proverimo da li korisnik sme da pristupi stranici

		next.ServeHTTP(w, r)

	})
}
