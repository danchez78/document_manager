package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response[T any] struct {
	Response T `json:"response"`
}

type Data[D any] struct {
	Data D `json:"data"`
}

type Error struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type ErrorResult struct {
	Error Error `json:"error"`
}

func ReturnResponse[T any](c echo.Context, resp T) error {
	return c.JSON(http.StatusOK, Response[T]{Response: resp})
}

func ReturnData[D any](c echo.Context, data D) error {
	return c.JSON(http.StatusOK, Data[D]{Data: data})
}

func ReturnFile(c echo.Context, contentType string, data []byte) error {
	return c.Blob(http.StatusOK, contentType, data)
}

func ReturnError(c echo.Context, statusCode int, err error) error {
	return c.JSON(http.StatusOK, ErrorResult{Error: Error{Code: statusCode, Text: err.Error()}})
}
