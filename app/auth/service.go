package auth

import (
	"alura-rest-base/errors"
	"alura-rest-base/types"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

	if validationError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); validationError != nil {
		return nil, errors.NewHttpError(http.StatusUnauthorized, "invalid credentials")
	}

	return &types.AuthenticateUserResponse{User: *user, Token: "token"}, nil
}

// GetUserIDFromToken implements types.IAuthService.
func (a *AuthService) GetUserIDFromToken(token string) (int, error) {
	panic("unimplemented")
}

func NewAuthService(userService types.IUserService) *AuthService {
	return &AuthService{userService: userService}
}
