package data

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserId   string
	Username string
	*jwt.StandardClaims
}
