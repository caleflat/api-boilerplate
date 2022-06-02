package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func Jwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
				return nil, fmt.Errorf("Invalid claims")
			}
			aud := "auth.kosherkit.io"
			checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAudience {
				return nil, fmt.Errorf(("invalid aud"))
			}
			iss := "api.kosherkit.io"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return nil, fmt.Errorf(("invalid iss"))
			}

			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			logrus.WithError(err).Error("Error parsing jwt token")
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	})
}
