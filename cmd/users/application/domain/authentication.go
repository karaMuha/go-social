package domain

import (
	"golang.org/x/crypto/bcrypt"
)

func Login(userID string, actualPassword string, passwordToCheck string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(actualPassword), []byte(passwordToCheck))
	if err != nil {
		return "", err
	}

	token, err := generateJwt(userID)
	if err != nil {
		return "", err
	}

	return token, nil
}
