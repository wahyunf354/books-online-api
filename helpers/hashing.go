package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password *string) (string, error) {
	// convert password string to byte
	passwordBytes := []byte(*password)

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPassword), err
}
