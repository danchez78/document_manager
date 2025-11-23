package usecases

import (
	"context"
	"document_manager/internal/application/domain"
)

type DeauthenticateUserHandler struct {
	authenticationService domain.AuthenticationService
	usersRepo             domain.UsersRepository
}

func NewDeauthenticateUserHandler(
	authenticationService domain.AuthenticationService,
	usersRepo domain.UsersRepository,
) *DeauthenticateUserHandler {
	if authenticationService == nil {
		panic("authentication service is nil")
	}
	if usersRepo == nil {
		panic("users repository is nil")
	}
	return &DeauthenticateUserHandler{
		authenticationService: authenticationService,
		usersRepo:             usersRepo,
	}
}

func (h *DeauthenticateUserHandler) Execute(ctx context.Context, tokenString string) error {
	user, err := h.authenticationService.ValidateToken(ctx, tokenString)
	if err != nil {
		return err
	}

	user.Deauthenticate()
	if err := h.usersRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
