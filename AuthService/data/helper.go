package data

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var SECRET = []byte("super-secret-AuthService-key")

func CreateJwt(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 1200).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return tokenString, nil
}
