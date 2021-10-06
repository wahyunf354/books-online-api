package payment_methods

import (
	"books_online_api/app/middlewares"
	"books_online_api/business/payment_methods"
	"books_online_api/controllers"
	"books_online_api/controllers/payment_methods/requests/create_payment_methods_requests"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PaymentMethodsController struct {
	PaymentMethodsUsecase payment_methods.Usecase
}

func NewPaymentMethodsController(usecase payment_methods.Usecase) *PaymentMethodsController {
	return &PaymentMethodsController{PaymentMethodsUsecase: usecase}
}

func (p PaymentMethodsController) CreatePaymentMethods(c echo.Context) error {

	cliams, err := middlewares.ExtractClaims(c)

	if cliams.Role != 3 {
		return controllers.NewErrorResponse(c, controllers.FORBIDDEN_USER)
	}

	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	var request create_payment_methods_requests.CreatePaymentMethodRequest

	err = c.Bind(&request)


	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	ctx := c.Request().Context()

	paymentMethod, err := p.PaymentMethodsUsecase.CreatePaymentMethod(ctx, request.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusOK, create_payment_methods_requests.FromDomain(paymentMethod))
}
