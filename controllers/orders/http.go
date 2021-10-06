package orders

import (
	"books_online_api/app/middlewares"
	"books_online_api/business/orders"
	"books_online_api/controllers"
	"books_online_api/controllers/orders/requests/create_order"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OrderController struct {
	OrderUsecase orders.Usecase
}

func NewOrdersController(orderUsecase orders.Usecase) *OrderController {
	return &OrderController{OrderUsecase: orderUsecase}
}

func (o OrderController) CreateOrder(c echo.Context) error {
	var request  create_order.OrderRequest

	err := c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	cliams,_ := middlewares.ExtractClaims(c)
	if cliams.Role < 1 {
		return controllers.NewErrorResponse(c, controllers.FORBIDDEN_USER)
	}

	request.UserId = cliams.UserId

	ctx := c.Request().Context()

	order, err := o.OrderUsecase.CreateOrder(ctx, request.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusOK, order)
}