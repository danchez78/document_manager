package error_handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"document_manager/internal/application/domain"
	"document_manager/internal/application/infrastructure/api/server"
)

func HandleError(c echo.Context, err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, domain.ErrEmptyToken) ||
		errors.Is(err, domain.ErrIncorrectPasswordProvided) ||
		errors.Is(err, domain.ErrTokenExpired) {
		return server.ReturnError(c, http.StatusUnauthorized, err)
	}

	if errors.Is(err, domain.ErrNoAccessToDoc) ||
		errors.Is(err, domain.ErrNoAccessToDeleteDoc) ||
		errors.Is(err, domain.ErrIncorrectAdminToken) {
		return server.ReturnError(c, http.StatusForbidden, err)
	}

	return server.ReturnError(c, http.StatusInternalServerError, err)
}
