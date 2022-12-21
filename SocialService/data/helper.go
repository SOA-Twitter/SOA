package data

import "github.com/golang-jwt/jwt/v4"

var SampleSecretKey = []byte("SecretYouShouldHide")

func GetFromClaims(token string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return SampleSecretKey, nil
	})
	return claims, err
}
