package user

import (
	"alura-rest-base/errors"
	"alura-rest-base/types"
	"net/http"
)

type UserService struct {
	storage types.IUserRepo
}

// GetUserByID implements types.IUserService.
func (u *UserService) GetUserByID(id int) (*types.User, *errors.HttpError) {
	user, err := u.storage.GetUserByID(id)
	if err != nil {
		return nil, errors.NewHttpError(http.StatusBadRequest, err.Error())
	}

	return user, nil
}

// AuthenticateUser implements types.IUserService.
func (u *UserService) GetUserByEmail(email string) (*types.User, *errors.HttpError) {
	user, err := u.storage.GetUserByEmail(email)
	if err != nil {
		return nil, errors.NewHttpError(http.StatusUnauthorized, "wrong credentials")
	}

	return user, nil
}

// CreateUser implements types.IUserService.
func (u *UserService) CreateUser(payload types.CreateUserPayload) (*types.User, *errors.HttpError) {
	user := &types.User{
		Email:     payload.Email,
		Password:  payload.Password, // TODO hash password
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	err := u.storage.CreateUser(user)
	if err != nil {
		return nil, errors.New500Error(err)
	}

	return user, nil
}

// DeleteUser implements types.IUserService.
func (u *UserService) DeleteUser(id int) *errors.HttpError {
	err := u.storage.DeleteUser(id)

	if err != nil {
		return errors.New500Error(err)
	}

	return nil
}

// Logout implements types.IUserService.
func (u *UserService) Logout() *errors.HttpError {
	panic("unimplemented")
}

func NewUserService(storage types.IUserRepo) *UserService {
	return &UserService{storage: storage}
}
