package middleware

import (
	"alura-go-base/types"
	"net/http"
)

func AuthMiddleware(userService types.IUserService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("unimplemented")
	})
}
