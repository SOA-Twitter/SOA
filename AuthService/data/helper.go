package data

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var SampleSecretKey = []byte("SecretYouShouldHide")

func CreateJwt(email string, role string) (string, error) {
	claims := &Claims{
		Role:  role,
		Email: email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 1200).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SampleSecretKey)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return tokenString, nil
}
func ValidateJwt(tokenString string) error {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SampleSecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return err
		}
		return err
	}
	// NOTE > read below
	if !token.Valid {
		return err
	}
	return nil
}

/*
	NOTE:  VALID = JWT Library functionality:

	func (c StandardClaims) Valid() error {
		vErr := new(ValidationError)
		now := TimeFunc().Unix()
		if !c.VerifyExpiresAt(now, false) {
			delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
			vErr.Inner = fmt.Errorf("%s by %s", ErrTokenExpired, delta)
			vErr.Errors |= ValidationErrorExpired
		}
*/
