package services

import (
	"alura-go-base/app/errors"
	"alura-go-base/types"
)

type UserService struct {
	storage types.IUserRepo
}

// GetUserByID implements types.IUserService.
func (u *UserService) GetUserByID(id int) errors.HttpError {
	panic("unimplemented")
}

// AuthenticateUser implements types.IUserService.
func (u *UserService) AuthenticateUser(payload types.AuthenticateUserPayload) (types.AuthenticateUserResponse, errors.HttpError) {
	panic("unimplemented")
}

// CreateUser implements types.IUserService.
func (u *UserService) CreateUser(payload types.CreateProductPayload) (*types.User, errors.HttpError) {
	panic("unimplemented")
}

// DeleteUser implements types.IUserService.
func (u *UserService) DeleteUser(id int) errors.HttpError {
	panic("unimplemented")
}

// Logout implements types.IUserService.
func (u *UserService) Logout() errors.HttpError {
	panic("unimplemented")
}

func NewUserService(storage types.IUserRepo) *UserService {
	return &UserService{storage: storage}
}
