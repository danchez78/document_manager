package error_handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"document_manager/internal/application/infrastructure/server/api/views"
	"document_manager/internal/application/usecases"
)

func HandleError(c echo.Context, err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, usecases.ErrEmptyToken) ||
		errors.Is(err, usecases.ErrIncorrectPasswordProvided) ||
		errors.Is(err, usecases.ErrTokenExpired) {
		return views.ReturnError(c, http.StatusUnauthorized, err)
	}

	if errors.Is(err, usecases.ErrNoAccessToDoc) ||
		errors.Is(err, usecases.ErrNoAccessToDeleteDoc) {
		return views.ReturnError(c, http.StatusForbidden, err)
	}

	return views.ReturnError(c, http.StatusInternalServerError, err)
}
