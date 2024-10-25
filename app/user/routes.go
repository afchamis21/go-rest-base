package user

import (
	"alura-rest-base/types"
	"net/http"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	userService types.IUserService
}

// RegisterHandlers implements types.IUserRouter.
func (u *UserRouter) RegisterHandlers(router *mux.Router) {
	subRouter := router.PathPrefix("/user").Subrouter()

	subRouter.HandleFunc("/logout", u.HandleLogout).Methods(http.MethodPost)
}

// HandleLogout implements types.IUserRouter.
func (u *UserRouter) HandleLogout(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewUserRouter(userService types.IUserService, router *mux.Router) *UserRouter {
	r := &UserRouter{userService: userService}
	r.RegisterHandlers(router)
	return r
}
