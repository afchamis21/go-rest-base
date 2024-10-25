package types

import (
	"alura-go-base/app/errors"
	"net/http"
)

type IAuthService interface {
	CreateUser(payload CreateUserPayload) (*User, *errors.HttpError)
	AuthenticateUser(payload AuthenticateUserPayload) (*AuthenticateUserResponse, *errors.HttpError)
	GetUserIDFromToken(token string) (int, error)
}

type IAuthRouter interface {
	Handler
	HandleRegisterUser(w http.ResponseWriter, r *http.Request)
	HandleAuthenticateUser(w http.ResponseWriter, r *http.Request)
}
