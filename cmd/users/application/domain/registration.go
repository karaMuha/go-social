package domain

import (
	"errors"
	"time"

	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

const REGISTRATION_TOKEN_LENGTH = 64

type Registration struct {
	ID                string `json:"id"`
	Username          string `json:"username" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"-" validate:"required"`
	RegistrationToken string
	Active            bool
	CreatedAt         time.Time `json:"created_at"`
}

// email and username must be case insensitively unique
// currently this is handled by the postgres implementation
// which makes the business logic depend on postgres
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

	registrationToken := randstr.String(REGISTRATION_TOKEN_LENGTH)
	user.RegistrationToken = registrationToken

	return &user, nil
}

func Activate(user *Registration, token string) error {
	if user.Active {
		return errors.New("user already active")
	}

	if len(user.RegistrationToken) != REGISTRATION_TOKEN_LENGTH {
		return errors.New("internal server error, token is funky")
	}

	if user.RegistrationToken != token {
		return errors.New("access denied")
	}

	return nil
}
