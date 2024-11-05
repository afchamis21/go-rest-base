package auth

import (
	"alura-rest-base/types"
	"alura-rest-base/util"
	"context"
	"net/http"
)

type contextKey string

const UserContextKey contextKey = "user"

func AuthMiddleware(authService types.IAuthService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			util.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		userId, err := authService.GetUserIDFromToken(token)
		if err != nil {
			util.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
