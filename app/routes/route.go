package routes

import (
	"books_online_api/controllers/users"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl * ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/login", cl.UserController.Register)
}