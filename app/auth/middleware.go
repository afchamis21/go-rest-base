package auth

import (
	"alura-go-base/types"
	"alura-go-base/util"
	"net/http"
)

func AuthMiddleware(userService types.IUserService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			util.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		if token != "token" { // TODO do the actual logic
			util.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
