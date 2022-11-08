package middleware

import (
	"TweeterMicro/AuthService/data"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

var SECRET = []byte("super-secret-AuthService-key")

func VerifyJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenString := c.Value
		claims := &data.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return SECRET, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

	})

}
