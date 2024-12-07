package domain

import (
	"golang.org/x/crypto/bcrypt"
)

func Login(userID string, hashedPassword string, passwordToCheck string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordToCheck))
	if err != nil {
		return "", err
	}

	token, err := GenerateJwt(userID)
	if err != nil {
		return "", err
	}

	return token, nil
}
