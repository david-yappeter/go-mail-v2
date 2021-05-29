package service

import (
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			http.Error(w, "Authorization Not Provided", http.StatusBadRequest)
			return
		}

		if !(AuthorizationCheck(auth)) {
			http.Error(w, "Bad Authorization", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
