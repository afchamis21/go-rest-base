package auth

import (
	"alura-rest-base/errors"
	"alura-rest-base/types"
	"net/http"
)

type AuthService struct {
	userService types.IUserService
}

// AuthenticateUser implements types.IAuthService.
func (a *AuthService) AuthenticateUser(payload types.AuthenticateUserPayload) (*types.AuthenticateUserResponse, *errors.HttpError) {
	user, err := a.userService.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, errors.NewHttpError(http.StatusUnauthorized, "invalid credentials")
	}

	// TODO finalize this
	if user.Password != payload.Password {
		return nil, errors.NewHttpError(http.StatusUnauthorized, "invalid credentials")
	}

	return &types.AuthenticateUserResponse{User: *user, Token: "token"}, nil
}

// CreateUser implements types.IAuthService.
func (a *AuthService) CreateUser(payload types.CreateUserPayload) (*types.User, *errors.HttpError) {
	user, err := a.userService.CreateUser(payload)

	if err != nil {
		return nil, err
	}

	return user, err
}

// GetUserIDFromToken implements types.IAuthService.
func (a *AuthService) GetUserIDFromToken(token string) (int, error) {
	panic("unimplemented")
}

func NewAuthService(userService types.IUserService) *AuthService {
	return &AuthService{userService: userService}
}
