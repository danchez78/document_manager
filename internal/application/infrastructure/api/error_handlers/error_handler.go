package error_handlers

import (
	"document_manager/internal/application/usecases"
	"document_manager/internal/common/server"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleError(c echo.Context, err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, usecases.ErrEmptyToken) ||
		errors.Is(err, usecases.ErrIncorrectPasswordProvided) ||
		errors.Is(err, usecases.ErrTokenExpired) {
		return server.ReturnError(c, http.StatusUnauthorized, err)
	}

	if errors.Is(err, usecases.ErrNoAccessToDoc) ||
		errors.Is(err, usecases.ErrNoAccessToDeleteDoc) {
		return server.ReturnError(c, http.StatusForbidden, err)
	}

	return server.ReturnError(c, http.StatusInternalServerError, err)
}
