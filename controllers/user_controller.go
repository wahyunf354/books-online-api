package controllers

import (
	"books_online_api/models/response"
	"books_online_api/models/users"
	"net/http"

	"github.com/labstack/echo"
)

func RegisterUser(c echo.Context) error {
	u := new(users.UserRegister)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Terjadi masalah pada request",
			Data:    err,
		})
	}

	return c.JSON(http.StatusOK, u)
}
