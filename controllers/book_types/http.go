package book_types

import (
	"books_online_api/app/middlewares"
	"books_online_api/business/book_types"
	"books_online_api/controllers"
	"books_online_api/controllers/book_types/requests"
	"books_online_api/controllers/book_types/responses"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookTypesController struct {
	BookTypesUsecase book_types.Usecase
}

func NewBookTypesController(bookTypeUsercase book_types.Usecase) *BookTypesController {
	return &BookTypesController{
		BookTypesUsecase: bookTypeUsercase,
	}
}

func (bookTypesController BookTypesController) CreateBookType(c echo.Context) error {
	bookTypeRequest := requests.BookTypeCreate{}
	err := c.Bind(&bookTypeRequest)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	cliams, err := middlewares.ExtractClaims(c)
	if cliams.Role != 3 {
		return controllers.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden user"))
	}

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	domain := bookTypeRequest.ToDomain()
	domain.Role = cliams.Role

	ctx := c.Request().Context()
	result, err := bookTypesController.BookTypesUsecase.CreateBookType(ctx, domain)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusCreated, responses.FromDomain(result))
}