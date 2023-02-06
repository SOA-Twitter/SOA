package data

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"io"
)

var SampleSecretKey = []byte("SecretYouShouldHide")

func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

func GetFromClaims(token string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return SampleSecretKey, nil
	})
	return claims, err
}
