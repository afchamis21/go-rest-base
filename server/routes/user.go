package routes

import (
	"alura-go-base/types"
	"net/http"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	userService types.IUserService
}

// RegisterHandlers implements types.IUserRouter.
func (u *UserRouter) RegisterHandlers(router *mux.Router) {
	panic("unimplemented")
}

// HandleAuthenticateUser implements types.IUserRouter.
func (u *UserRouter) HandleAuthenticateUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// HandleDeleteUser implements types.IUserRouter.
func (u *UserRouter) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// HandleLogout implements types.IUserRouter.
func (u *UserRouter) HandleLogout(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// HandleRegisterUser implements types.IUserRouter.
func (u *UserRouter) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewUserRouter(userService types.IUserService, router *mux.Router) *UserRouter {
	r := &UserRouter{userService: userService}
	r.RegisterHandlers(router)
	return r
}
