package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetUserController(c echo.Context) error {
	users := User{"Wahyu", "wahyo@hh.com", "malang"}

	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    users,
	})
}

func main() {
	e := echo.New()

	e.GET("v1/users", GetUserController)

	e.Start(":8000")
}
