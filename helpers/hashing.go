package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) (string, error) {
	// convert password string to byte
	passwordBytes := []byte(*password)

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPassword), err
}

func ComparePassword(hash string, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))
	if err != nil {
		return  false
	}
	return true
}
