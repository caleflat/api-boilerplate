package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Debug("Request received")

		next.ServeHTTP(w, r)
	})
}
