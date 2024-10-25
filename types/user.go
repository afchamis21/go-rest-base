package types

import (
	"alura-rest-base/errors"
	"net/http"
)

type User struct {
	ID        int    `json:"-"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreateUserPayload struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type AuthenticateUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthenticateUserResponse struct {
	User  User
	Token string
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
	Logout() *errors.HttpError
}

type IUserRouter interface {
	IHandler
	HandleLogout(w http.ResponseWriter, r *http.Request)
}
