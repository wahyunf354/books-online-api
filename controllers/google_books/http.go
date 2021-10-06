package google_books

import (
	"books_online_api/business/google_books"
	"books_online_api/controllers"
	"books_online_api/controllers/google_books/requests"
	"books_online_api/controllers/google_books/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type GoogleBooksController struct {
	GoogleBookUsecase google_books.Usecase
}

func NewGoogleBooksController(googleBookUsecase google_books.Usecase) *GoogleBooksController {
	return &GoogleBooksController{
		GoogleBookUsecase: googleBookUsecase,
	}
}

func (gbController GoogleBooksController) SearchBooks(c echo.Context) error {
	bookRequestParams := requests.GoogleBookRequest{}
	bookRequestParams.Keyword = c.QueryParam("q")
	bookRequestParams.StartIndex, _ = strconv.Atoi(c.QueryParam("startIndex"))
	bookRequestParams.MaxResults, _ = strconv.Atoi(c.QueryParam("maxResult"))

	googleBookDomain := bookRequestParams.ToDomain()

	ctx := c.Request().Context()
	resultGoogleBooks, err := gbController.GoogleBookUsecase.SearchBooks(ctx, googleBookDomain.Keyword, googleBookDomain.StartIndex, googleBookDomain.MaxResults)
	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}
	 return controllers.NewSuccessResponse(c, http.StatusOK, responses.FormListDomain(resultGoogleBooks))
}