package data

import (
	"TweeterMicro/auth/model"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var SECRET = []byte("super-secret-auth-key")

func CreateJwt(email string, password string) (string, error) {
	claims := &model.Claims{
		Email:    email,
		Password: password,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 1200).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return tokenString, nil
}
