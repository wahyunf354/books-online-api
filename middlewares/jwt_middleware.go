package middlewares

import (
	"books_online_api/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(email string, firstName string, lastName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["firstName"] = firstName
	claims["lastName"] = lastName
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}
