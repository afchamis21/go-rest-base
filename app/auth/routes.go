package auth

import (
	"alura-rest-base/types"
	"alura-rest-base/util"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthRouter struct {
	authService types.IAuthService
	userService types.IUserService
}

// HandleAuthenticateUser implements types.IAuthRouter.
func (a *AuthRouter) HandleAuthenticateUser(w http.ResponseWriter, r *http.Request) {
	payload, e := util.ReadJson[types.AuthenticateUserPayload](r)
	if e != nil {
		util.WriteError(w, http.StatusBadRequest, e.Error())
		return
	}

	res, err := a.authService.AuthenticateUser(payload)
	if err != nil {
		util.WriteError(w, err.StatusCd, err.Message)
		return
	}

	util.WriteJson(w, http.StatusOK, res)
}

// HandleRegisterUser implements types.IAuthRouter.
func (a *AuthRouter) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	payload, e := util.ReadJson[types.CreateUserPayload](r)
	if e != nil {
		util.WriteError(w, http.StatusBadRequest, e.Error())
		return
	}

	res, err := a.userService.CreateUser(payload)
	if err != nil {
		util.WriteError(w, err.StatusCd, err.Message)
		return
	}

	util.WriteJson(w, http.StatusCreated, res)
}

// RegisterHandlers implements types.IAuthRouter.
func (a *AuthRouter) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/login", a.HandleAuthenticateUser).Methods(http.MethodPost)
	router.HandleFunc("/register", a.HandleRegisterUser).Methods(http.MethodPost)
}

func NewAuthRouter(
	router *mux.Router,
	authService types.IAuthService,
	userService types.IUserService,
) *AuthRouter {
	r := &AuthRouter{
		authService: authService,
		userService: userService,
	}

	r.RegisterHandlers(router)
	return r
}
