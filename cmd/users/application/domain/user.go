package domain

import (
	"time"

	"github.com/karaMuha/go-social/users/application/utils"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"-" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func RegisterUser(username, email, password string) (*User, error) {
	user := User{
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	err := utils.Validator.Check(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
