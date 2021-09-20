package routes

import (
	"books_online_api/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("/api/v1")
	ev1.POST("/users/register", controllers.RegisterUser)

	ev1.POST("/booktypes", controllers.CreateBookTypes)

	return e
}
