package controllers

import (
	"books_online_api/configs"
	"books_online_api/helpers"
	"books_online_api/models/books"
	"books_online_api/models/response"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func UploadFileBook(fileBook *multipart.FileHeader) (result string, err error, message string) {
	// Save book
	srcFileBooks, err := fileBook.Open()
	if err != nil {
		return "", err, "Internal Server Error"
	}
	defer srcFileBooks.Close()
	fileNameBook := fileBook.Filename
	splitFileNameBook := strings.Split(fileNameBook, ".")
	// Check page extension page book
	extentionFileBook := splitFileNameBook[len(splitFileNameBook)-1]
	validExtensionFileBook := []string{"pdf"}
	if !helpers.CheckItemInSlice(validExtensionFileBook, extentionFileBook) {
		return "", echo.ErrInternalServerError, "Input file book no pdf"
	}
	newFileNameBook := time.Now().Format(time.RFC850) + "_" + helpers.RandomString(10) + "_" + splitFileNameBook[0] + "." + extentionFileBook
	dstFileBook, err := os.Create("./public/files/file_books/" + newFileNameBook)

	if err != nil {
		return "", err, "Internal server error"
	}
	defer dstFileBook.Close()
	if _, err := io.Copy(dstFileBook, srcFileBooks); err != nil {
		return "", err, "Internal server error"
	}
	return newFileNameBook, nil, "Success upload file book"
}

func UploadImages(images []*multipart.FileHeader) ([]string, error, string) {
	result := []string{}
	for _, image := range images {
		src, err := image.Open()
		if err != nil {
			return result, err, "Internet Server Error"
		}
		defer src.Close()

		fileName := image.Filename
		splitFileName := strings.Split(fileName, ".")

		// Check image extention
		extention := splitFileName[len(splitFileName)-1]
		validExtensionImage := []string{"jpg", "jpeg", "png"}
		if !helpers.CheckItemInSlice(validExtensionImage, extention) {
			return result, nil, "No input image"
		}
		newFileName := time.Now().Format(time.RFC850) + "_" + helpers.RandomString(10) + "_" + splitFileName[0] + "." + extention

		dst, err := os.Create("./public/img/img_books/" + newFileName)
		if err != nil {
			return result, err, "Internal server error"
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return result, err, "Internal server error"
		}

		result = append(result, newFileName)
	}
	return result, nil, "Success Upload Gambar"
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

	// Melakukan validasi field
	if book.BookTypeId == 0 || book.CategoryId == 0 || book.UserId == 0 || book.Description == "" || book.PageCount == 0 || book.Price == 0 || book.Title == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Some fields are still empty",
			Data:    nil,
		})
	}

	// MENYIMPAN BUKU
	fileBook, err := c.FormFile("file_book")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "empty book file",
			Data:    err,
		})
	}
	newFileBook, err, message := UploadFileBook(fileBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: message,
			Data:    err,
		})
	}

	// MENYIMPAN DATA IMAGE
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error sini",
			Data:    err,
		})
	}
	images := form.File["image_books"]
	// upload file image
	nameImagesNew, err, message := UploadImages(images)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: message,
			Data:    err,
		})
	}
	// Melakukan penyimpanan ke DB
	// Menyimpan data ke table book
	bookNew := books.Book{}
	bookNew.UserId = book.UserId
	bookNew.BookTypeId = book.BookTypeId
	bookNew.CategoryId = book.BookTypeId
	bookNew.Title = book.Title
	bookNew.Price = book.Price

	bookResult := configs.DB.Create(&bookNew)
	if bookResult.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "book failed to enter into db",
			Data:    err,
		})
	}

	// Menyimpan data ke table book description
	bookDescriptionNew := books.BookDetail{}
	bookDescriptionNew.PageCount = book.PageCount
	bookDescriptionNew.UrlBook = newFileBook
	bookDescriptionNew.Description = book.Description
	bookDescriptionNew.BookId = bookNew.ID
	bookDetailResult := configs.DB.Create(&bookDescriptionNew)

	if bookDetailResult.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "book details failed to enter into db",
			Data:    err,
		})
	}

	// Menyimpan Images
	for _, url := range nameImagesNew {
		image := books.ImageBooks{}
		image.BookId = bookNew.ID
		image.Url = url
		imageResutl := configs.DB.Create(&image)
		if imageResutl.Error != nil {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "image failed to enter into db",
				Data:    err,
			})
		}
	}

	return c.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Success create book",
		Data:    bookNew,
	})
}
