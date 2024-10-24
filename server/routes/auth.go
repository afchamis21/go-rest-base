package routes

import (
	"alura-go-base/types"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthRouter struct {
	authService types.IAuthService
}

// HandleAuthenticateUser implements types.IAuthRouter.
func (a *AuthRouter) HandleAuthenticateUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// HandleRegisterUser implements types.IAuthRouter.
func (a *AuthRouter) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// RegisterHandlers implements types.IAuthRouter.
func (a *AuthRouter) RegisterHandlers(router *mux.Router) {
	panic("unimplemented")
}

func NewAuthRouter(authService types.IAuthService, router *mux.Router) *AuthRouter {
	r := &AuthRouter{authService: authService}
	r.RegisterHandlers(router)

	return r
}
