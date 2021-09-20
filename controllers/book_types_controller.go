package controllers

import (
	"books_online_api/configs"
	"books_online_api/models/books"
	"books_online_api/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBookTypes(c echo.Context) error {
	bookType := new(books.BookTypeRequest)

	// Melakukan Bind Data
	if err := c.Bind(bookType); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Terjadi masalah pada request",
			Data:    err,
		})
	}

	// Validasi jika ada data yang kosong
	if bookType.Name == "" || bookType.Unit == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Field name, unit tidak boleh kosong",
			Data:    nil,
		})
	}

	// Input data ke DB
	bookTypeNew := books.BookType{}
	bookTypeNew.Name = bookType.Name
	bookTypeNew.Unit = bookType.Unit

	result := configs.DB.Create(&bookTypeNew)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Terjadi masalah saat input data ke DB",
			Data:    result.Error,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil Menambahkan Book Types",
		Data:    bookTypeNew,
	})
}
