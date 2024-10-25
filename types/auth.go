package types

import (
	"alura-rest-base/errors"
	"net/http"
)

type IAuthService interface {
	AuthenticateUser(payload AuthenticateUserPayload) (*AuthenticateUserResponse, *errors.HttpError)
	GetUserIDFromToken(token string) (int, error)
}

type IAuthRouter interface {
	IHandler
	HandleRegisterUser(w http.ResponseWriter, r *http.Request)
	HandleAuthenticateUser(w http.ResponseWriter, r *http.Request)
}

type AuthenticateUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthenticateUserResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
