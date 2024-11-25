package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Registration struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"-" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func Signup(username, email, password string) (*Registration, error) {
	user := Registration{
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	err := validate.Struct(&user)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	return &user, nil
}
