package books

import (
	"books_online_api/app/middlewares"
	"books_online_api/business/books"
	"books_online_api/controllers"
	"books_online_api/controllers/books/requests"
	"books_online_api/controllers/books/responses/create_books"
	"books_online_api/controllers/books/responses/get_books"
	"books_online_api/controllers/books/responses/get_one_book"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
		return controllers.NewErrorResponse(c, controllers.FORBIDDEN_USER)
	}

	bookRequest := requests.BookRequest{}

	err = c.Bind(&bookRequest)
	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	bookRequest.FileBook, err = c.FormFile("file_book")
	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	bookRequest.ImagesBook = form.File["images"]
	bookRequest.UserId = cliams.UserId

	ctx := c.Request().Context()

	booksDomain, err := bookController.BooksUsecase.CreateBook(ctx, bookRequest.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, err)
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
		return controllers.NewErrorResponse(c, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusOK, get_books.FromListDomain(result))
}

func (booksController BooksController) GetOneBook(c echo.Context)  error {
	var request requests.GetOneBookRequest
	request.Id, _ = strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()

	book, err := booksController.BooksUsecase.GetOneBook(ctx, request.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusOK, get_one_book.FromDomain(book))
}