package users

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPasswordWithSalt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("[warning] unable to generate hash of password with length", len(password))
		return "", err
	}
	return string(bytes), nil
}

func ValidatePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
