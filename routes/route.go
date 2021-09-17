package routes

import (
	"books_online_api/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("/api/v1")
	ev1.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    200,
			Message: "Success Running",
			Data:    "Tidak Ada",
		})
	})
	return e
}
