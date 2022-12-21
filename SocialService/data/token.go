package data

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Role     string
	Email    string
	Username string
	*jwt.StandardClaims
}
