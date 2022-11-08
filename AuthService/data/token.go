package data

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string
	Password string
	*jwt.StandardClaims
}
