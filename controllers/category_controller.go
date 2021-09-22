package controllers

import (
	"books_online_api/configs"
	"books_online_api/models/books"
	"books_online_api/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCategory(c echo.Context) error {

	category := new(books.CategoryRequest)

	// Melakukan Bind Data
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Terjadi masalah pada request",
			Data:    err,
		})
	}

	// Validasi title category
	if category.Title == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Field title tidak boleh kosong",
			Data:    nil,
		})
	}

	newCategory := books.Category{}
	newCategory.Title = category.Title

	result := configs.DB.Create(&newCategory)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Terjadi masalah saat memasukan data pada DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Success create category book",
		Data:    newCategory,
	})
}
