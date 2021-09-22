package routes

import (
	"books_online_api/controllers"
	m "books_online_api/middlewares"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	// Middleware
	m.RemoveSlashMiddleware(e)
	m.LogMiddleware(e)

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
