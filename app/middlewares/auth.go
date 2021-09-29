package middlewares

import (
	"books_online_api/controllers"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

type JwtCustomClaims struct {
	UserId int `json:"user_id"`
	Role int8 `json:"role"`
	jwt.StandardClaims
}

type ConfigJwt struct {
	SecretJwt string
	ExpiredDuration int
}

func (configJWT *ConfigJwt) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims: &JwtCustomClaims{},
		SigningKey: []byte(configJWT.SecretJwt),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(err error, context echo.Context) error {
			return controllers.NewErrorResponse(context, http.StatusForbidden, err)
		}),
	}
}


func ExtractClaims(c echo.Context) (*JwtCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims, nil
}

func (configJWT *ConfigJwt) GenerateTokenJWT(userId int, role int8) (string, error) {
	claims := JwtCustomClaims{
		userId,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(configJWT.ExpiredDuration))).Unix(),
		},
	}

	//create token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(configJWT.SecretJwt))

	if err != nil {
		return "", err
	}
	return token, err
}