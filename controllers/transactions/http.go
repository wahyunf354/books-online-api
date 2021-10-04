package transactions

import (
	"books_online_api/app/middlewares"
	"books_online_api/business/transactions"
	"books_online_api/controllers"
	"books_online_api/controllers/transactions/requests"
	"books_online_api/controllers/transactions/responses/update_transaction"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TransactionsController struct {
	TransactionsUsecase transactions.Usecase
}

func NewTransactionsControllers(usecase transactions.Usecase) *TransactionsController {
	return &TransactionsController{TransactionsUsecase: usecase}
}

func (t TransactionsController) CreateTransactions(c echo.Context) error {
	var request requests.CreateTransactionRequest

	err := c.Bind(&request)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	cliams, err := middlewares.ExtractClaims(c)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusForbidden, err)
	}

	request.UserId = cliams.UserId

	ctx := c.Request().Context()
	result, err := t.TransactionsUsecase.CreateTransactions(ctx, request.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusOK, update_transaction.FromDomain(result))
}

func (t TransactionsController) UpdateStatusTransactions(c echo.Context) error {
	var request requests.UpdateStatusTransactionRequest
	var err error
	var transaction transactions.Domain

	cliams, err := middlewares.ExtractClaims(c)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	if cliams.Role != 3 {
		return controllers.NewErrorResponse(c, http.StatusForbidden, errors.New("user forbidden"))
	}

	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	transaction, err = t.TransactionsUsecase.UpdateStatusTransaction(ctx, request.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusOK, update_transaction.FromDomain(transaction))
}