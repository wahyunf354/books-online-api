package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, status int, data interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Success"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, err error) error {
	status := CheckStatus(err)
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = err.Error()
	response.Data = nil
	return c.JSON(status, response)
}

func CheckStatus(err error) int {
	if err == EMPTY_TITLE {
		return http.StatusBadRequest
	}
	if err == EMPTY_DESCRIPTION {
		return http.StatusBadRequest
	}
	if err == FORBIDDEN_USER {
		return http.StatusForbidden
	}
	if err == ID_EMPTY {
		return http.StatusBadRequest
	}
	if err == NAME_BOOK_TYPE_EMPTY {
		return http.StatusBadRequest
	}
	if err == UNIT_BOOK_TYPE_EMPTY {
		return http.StatusBadRequest
	}
	if err == RECORD_NOT_FOUND {
		return http.StatusNoContent
	}
	if err == KEYWORD_EMPTY {
		return http.StatusBadRequest
	}
	if err == NAME_EMPTY {
		return http.StatusBadRequest
	}
	if err == ADDRESS_EMPTY {
		return http.StatusBadRequest
	}
	if err == AUTHOR_EMPTY {
		return http.StatusBadRequest
	}
	if err == ORDER_ID_EMPTY {
		return http.StatusBadRequest
	}
	if err == EMPTY_EMAIL ||
		err == INVALID_EMAIL ||
		err == EMPTY_FIRST_NAME ||
		err == EMPTY_LAST_NAME ||
		err == EMPTY_PASSWORD ||
		err == EMPTY_CONF_PASSWORD {
		return http.StatusBadRequest
	}
	if err == WRONG_PASSWORD {
		return http.StatusForbidden
	}
	return http.StatusInternalServerError
}
