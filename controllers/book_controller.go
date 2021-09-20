package controllers

import (
	"books_online_api/models/books"
	"books_online_api/models/response"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func CheckItemInSlice(arr []string, item string) bool {
	for _, e := range arr {
		if strings.ToLower(e) == strings.ToLower(item) {
			return true
		}
	}
	return false
}

func CreateBook(c echo.Context) error {
	book := new(books.BookRequest)

	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "There is a problem with the request",
			Data:    nil,
		})
	}
	// TODO: Melakukan validasi
	if book.BookTypeId == 0 || book.CategoryId == 0 || book.UserId == 0 || book.Description == "" || book.PageCount == 0 || book.Price == 0 || book.Title == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Some fields are still empty",
			Data:    nil,
		})
	}

	// menyimpan buku
	fileBook, err := c.FormFile("file_book")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    nil,
		})
	}
	fmt.Println(fileBook)
	// TODO: Save book

	// menyimpan data Image
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    err,
		})
	}

	for _, image := range form.File["images"] {
		src, err := image.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
				Data:    err,
			})
		}
		defer src.Close()

		fileName := image.Filename
		splitFileName := strings.Split(fileName, ".")

		// Check image extention
		extention := splitFileName[len(splitFileName)-1]
		if !CheckItemInSlice([]string{"jpg", "jpeg", "png"}, extention) {
			return c.JSON(http.StatusBadRequest, response.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "Input not images",
				Data:    nil,
			})
		}

		newFileName := time.Now().Format(time.RFC850) + "_" + RandomString(10) + "_" + splitFileName[0] + "." + extention

		dst, err := os.Create("./public/img/img_books/" + newFileName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
				Data:    err,
			})
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
				Data:    err,
			})
		}

		book.ImageUrls = append(book.ImageUrls, newFileName)
	}

	// TODO: Melakukan penyimpanan DB
	bookNew := books.Book{}
	bookNew.UserId = book.UserId
	bookNew.BookTypeId = book.BookTypeId
	bookNew.CategoryId = book.BookTypeId
	bookNew.Title = book.Title
	bookNew.Price = book.Price
	bookDescriptionNew := books.BookDetail{}
	bookDescriptionNew.PageCount = book.PageCount

	// configs.DB.Transaction(func(tx *gorm.DB) error {

	// })

	return c.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Success create book",
		Data:    book,
	})
}
