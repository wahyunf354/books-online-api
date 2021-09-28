package books

import (
	"books_online_api/business/books"
	"books_online_api/controllers"
	"books_online_api/controllers/books/requests"
	"books_online_api/controllers/books/responses/create_books"
	"github.com/labstack/echo/v4"
	"net/http"
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
	bookRequest := requests.BookRequest{}
	var err error

	err = c.Bind(&bookRequest)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	bookRequest.FileBook, err = c.FormFile("file_book")

	ctx := c.Request().Context()

	booksDomain, err := bookController.BooksUsecase.CreateBook(ctx, bookRequest.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusOK, create_books.FromDomain(booksDomain))
}

