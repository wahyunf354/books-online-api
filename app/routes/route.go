package routes

import (
	"books_online_api/controllers/users"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl * ControllerList) RouteRegister(e *echo.Echo) {
	ev1 := e.Group("api/v1/")
	ev1.POST("users/register", cl.UserController.Register)
}