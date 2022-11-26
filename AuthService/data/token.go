package data

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Email string
	*jwt.StandardClaims
}
