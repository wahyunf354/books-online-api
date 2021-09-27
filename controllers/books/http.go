package books

import (
	"books_online_api/business/books"
	"context"
	"github.com/labstack/echo/v4"
)

type BooksController struct {
	BooksUsecase books.Usecase
}

func NewBooksController(booksUsecase books.Usecase) *BooksController {
	return &BooksController{
		BooksUsecase: booksUsecase,
	}
}

func (bookController BooksController) CreateBook(c echo.Context) error {

}

