package main

import (
	"books_online_api/configs"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	configs.InitDB()

	e := echo.New()
	ev1 := e.Group("/api/v1")
	ev1.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, BaseResponse{
			Code:    200,
			Message: "Success Running",
			Data:    "Tidak Ada",
		})
	})

	e.Logger.Fatal(e.Start(":8001"))

}
