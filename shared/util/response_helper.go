package util

import (
	echo "github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	Title         string `json:"title,omitempty"`
	SystemMessage string `json:"system_message"`
}

type JSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponse(c echo.Context, statusCode int, msg string, data interface{}) error {
	return c.JSON(statusCode, JSONResponse{
		Code:    statusCode,
		Message: msg,
		Data:    data,
	})
}

func SetErrorResponse(c echo.Context, errorResponse ErrorResponse) error {
	return c.JSON(errorResponse.Code, errorResponse)
}
