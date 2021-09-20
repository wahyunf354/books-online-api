package routes

import (
	"books_online_api/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("/api/v1")
	// Users
	ev1.POST("/users/register", controllers.RegisterUser)

	// Book Type
	ev1.POST("/booktypes", controllers.CreateBookType)

	// Category Book
	ev1.POST("/categories", controllers.CreateCategory)

	return e
}
