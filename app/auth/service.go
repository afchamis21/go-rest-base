package auth

import (
	"alura-rest-base/config"
	"alura-rest-base/errors"
	"alura-rest-base/types"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var tokenIssuer = "@afchamis/go-rest-base"

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

	token, e := a.generateJWT(user.ID)
	if e != nil {
		return nil, errors.New500Error(e)
	}

	return &types.AuthenticateUserResponse{User: *user, Token: token}, nil
}

// GetUserIDFromToken implements types.IAuthService.
func (a *AuthService) GetUserIDFromToken(tokenStr string) (int, *errors.HttpError) {
	t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := []byte(config.Envs.AuthTokenSecret)
		return secret, nil
	})

	if err != nil {
		return 0, errors.NewHttpError(http.StatusUnauthorized, "invalid token")
	}

	issuer, err := t.Claims.GetIssuer()
	if err != nil {
		return 0, errors.New500Error(err)
	} else if issuer != tokenIssuer {
		return 0, errors.NewHttpError(http.StatusUnauthorized, "invalid token")
	}

	subject, err := t.Claims.GetSubject()
	if err != nil {
		return 0, errors.New500Error(err)
	}

	userId, err := strconv.Atoi(subject)
	if err != nil {
		return 0, errors.New500Error(err)
	}

	return userId, nil
}

func (a *AuthService) generateJWT(userId int) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    tokenIssuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		Subject:   fmt.Sprint(userId),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(config.Envs.AuthTokenSecret)

	return token.SignedString(secret)
}

func NewAuthService(userService types.IUserService) *AuthService {
	return &AuthService{userService: userService}
}
