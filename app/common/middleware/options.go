package middleware

import (
	"net/http"
)

func OptionsMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}