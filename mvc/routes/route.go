package routes

import (
	"books_online_api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	ev1 := e.Group("/api/v1")
	// Users
	ev1.POST("/users/register", controllers.RegisterUser)
	ev1.POST("/users/login", controllers.LoginUser)

	// Book Type
	ev1.POST("/booktypes", controllers.CreateBookType)

	// Category Book
	ev1.POST("/categories", controllers.CreateCategory)

	// Books
	ev1.POST("/books", controllers.CreateBook)
	return e
}
