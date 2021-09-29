package books

import (
	"books_online_api/app/middlewares"
	"books_online_api/business/books"
	"books_online_api/controllers"
	"books_online_api/controllers/books/requests"
	"books_online_api/controllers/books/responses/create_books"
	"errors"
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
	var err error

	cliams, err := middlewares.ExtractClaims(c)
	if cliams.Role < 2 {
		return controllers.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden user"))
	}

	bookRequest := requests.BookRequest{}

	err = c.Bind(&bookRequest)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	bookRequest.FileBook, err = c.FormFile("file_book")
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	bookRequest.ImagesBook = form.File["images"]

	ctx := c.Request().Context()

	booksDomain, err := bookController.BooksUsecase.CreateBook(ctx, bookRequest.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	// TODO: RESPONSE URL FILE NYA DIBUAT FULL DENGAN HOST
	return controllers.NewSuccessResponse(c, http.StatusOK, create_books.FromDomain(booksDomain))
}

func (booksController BooksController) GetBooks(c echo.Context) error {
	var request requests.GetBooksRequest

	request.Keyword = c.QueryParam("q")

	ctx := c.Request().Context()

	result, err := booksController.BooksUsecase.GetBooks(ctx, request.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusOK, result)
}