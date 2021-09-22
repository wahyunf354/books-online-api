package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	UserId int `json:"userid"`
	jwt.StandardClaims
}

func CreateToken(id int) (string, error) {
	claims := &JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 1).Unix(),
		},
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	// TODO: Masih error untuk generate JWT
	jwtToken, err := token.SignedString([]byte("hallo"))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
