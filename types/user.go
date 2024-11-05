package types

import (
	"alura-rest-base/errors"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreateUserPayload struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type IUserRepo interface {
	CreateUser(user *User) error
	DeleteUser(id int) error
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type IUserService interface {
	CreateUser(payload CreateUserPayload) (*User, *errors.HttpError)
	DeleteUser(id int) *errors.HttpError
	GetUserByID(id int) (*User, *errors.HttpError)
	GetUserByEmail(email string) (*User, *errors.HttpError)
}
