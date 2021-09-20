package controllers

import (
	"books_online_api/models/books"
	"books_online_api/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	book := new(books.BookRequest)

	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Terjadi masalah pada request",
			Data:    nil,
		})
	}
	// TODO: Melakukan validasi

	// TODO: menyimpan data Image

	// TODO: Melakukan penyimpanan DB

	return c.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Success membuat book",
		Data:    book,
	})
}
