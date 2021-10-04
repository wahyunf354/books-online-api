package routes

import (
	"books_online_api/controllers/book_types"
	"books_online_api/controllers/books"
	"books_online_api/controllers/google_books"
	"books_online_api/controllers/orders"
	"books_online_api/controllers/payment_methods"
	"books_online_api/controllers/transactions"
	"books_online_api/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware middleware.JWTConfig
	UserController users.UserController
	GoogleBooksController google_books.GoogleBooksController
	BookTypeController book_types.BookTypesController
	BooksController books.BooksController
	OrdersController orders.OrderController
	PaymentMethodsController payment_methods.PaymentMethodsController
	TransactionsController transactions.TransactionsController
}


func (cl * ControllerList) RouteRegister(e *echo.Echo) {
	ev1 := e.Group("api/v1/")
	ev1.POST("users/register", cl.UserController.Register)
	ev1.POST("users/login", cl.UserController.Login)

	withJwt := ev1.Group("books")
	withJwt.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	withJwt.GET("/googlebook", cl.GoogleBooksController.SearchBooks)
	withJwt.POST("/booktype", cl.BookTypeController.CreateBookType)
	withJwt.POST("", cl.BooksController.CreateBook)
	withJwt.GET("", cl.BooksController.GetBooks)
	withJwt.GET("/:id", cl.BooksController.GetOneBook)

	orderWithJwt := ev1.Group("orders")
	orderWithJwt.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	orderWithJwt.POST("", cl.OrdersController.CreateOrder)

	paymentMethodWithJwt := ev1.Group("payments")
	paymentMethodWithJwt.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	paymentMethodWithJwt.POST("", cl.PaymentMethodsController.CreatePaymentMethods)

	transactionsJWT := ev1.Group("transactions")
	transactionsJWT.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	transactionsJWT.POST("", cl.TransactionsController.CreateTransactions)
}
