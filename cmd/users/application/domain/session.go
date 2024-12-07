package domain

import (
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(hashedPassword string, passwordToCheck string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordToCheck))
}
