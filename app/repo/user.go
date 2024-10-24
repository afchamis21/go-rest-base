package repo

import (
	"alura-go-base/types"
	"database/sql"
	"errors"
)

type UserRepo struct {
	db *sql.DB
}

// CreateUser implements types.IUserRepo.
func (u *UserRepo) CreateUser(user *types.User) error {
	err := u.db.QueryRow("INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING ID",
		user.Email, user.Password, user.FirstName, user.LastName,
	).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteUser implements types.IUserRepo.
func (u *UserRepo) DeleteUser(id int) error {
	_, err := u.db.Exec("DELETE FROM users u WHERE u.ID = $1", id)

	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail implements types.IUserRepo.
func (u *UserRepo) GetUserByEmail(email string) (*types.User, error) {
	rows, err := u.db.Query("SELECT * FROM users u WHERE u.EMAIL=$1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user, err := mapRowToUser(rows)

		if err != nil {
			return nil, err
		}

		if user.ID == 0 {
			return nil, errors.New("user not found")
		}

		return user, nil
	}

	return nil, errors.New("user not found")
}

// GetUserByID implements types.IUserRepo.
func (u *UserRepo) GetUserByID(id int) (*types.User, error) {
	rows, err := u.db.Query("SELECT * FROM users u WHERE u.ID=$1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user, err := mapRowToUser(rows)

		if err != nil {
			return nil, err
		}

		if user.ID == 0 {
			return nil, errors.New("user not found")
		}

		return user, nil
	}

	return nil, errors.New("user not found")
}

func mapRowToUser(rows *sql.Rows) (*types.User, error) {
	user := types.User{}
	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}
