package google_books

import (
	"books_online_api/business/google_books"
	"books_online_api/controllers"
	"books_online_api/controllers/google_books/requests"
	"books_online_api/controllers/google_books/responses"
	"github.com/labstack/echo/v4"
	"net/http"
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

	ctx := c.Request().Context()
	resultGoogleBooks, err := gbController.GoogleBookUsecase.SearchBooks(ctx, bookRequestParams.ToDomain().Keyword)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	 return controllers.NewSuccessResponse(c, http.StatusOK, responses.FormListDomain(resultGoogleBooks))
}