package routes

import (
	"books_online_api/controllers/google_books"
	"books_online_api/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware middleware.JWTConfig
	UserController users.UserController
	GoogleBooksController google_books.GoogleBooksController
}

func (cl * ControllerList) RouteRegister(e *echo.Echo) {
	ev1 := e.Group("api/v1/")
	ev1.POST("users/register", cl.UserController.Register)
	ev1.POST("users/login", cl.UserController.Login)

	ev1.GET("books/googlebooks", cl.GoogleBooksController.SearchBooks)
}