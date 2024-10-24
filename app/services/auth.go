package services

import (
	"alura-go-base/app/errors"
	"alura-go-base/types"
)

type AuthService struct {
	userService UserService
}

// AuthenticateUser implements types.IAuthService.
func (a *AuthService) AuthenticateUser(payload types.AuthenticateUserPayload) (types.AuthenticateUserResponse, errors.HttpError) {
	panic("unimplemented")
}

// CreateUser implements types.IAuthService.
func (a *AuthService) CreateUser(payload types.CreateProductPayload) (*types.User, errors.HttpError) {
	panic("unimplemented")
}

// GetUserIDFromToken implements types.IAuthService.
func (a *AuthService) GetUserIDFromToken(token string) (int, error) {
	panic("unimplemented")
}

func NewAuthService(userService UserService) *AuthService {
	return &AuthService{userService: userService}
}
