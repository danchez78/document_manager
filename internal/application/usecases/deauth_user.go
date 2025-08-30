package usecases

import (
	"context"

	"document_manager/internal/application/domain"
)

type DeauthUserHandler struct {
	usersRepo domain.UsersRepository
	tm        domain.TokenManager
}

func NewDeauthUserHandler(usersRepo domain.UsersRepository, tm domain.TokenManager) *DeauthUserHandler {
	if usersRepo == nil {
		panic("users repository is nil")
	}

	if tm == nil {
		panic("token manager is nil")
	}

	return &DeauthUserHandler{usersRepo: usersRepo, tm: tm}
}

func (h *DeauthUserHandler) Execute(ctx context.Context, tokenString string) error {
	if tokenString == "" {
		return ErrEmptyToken
	}
	token, err := h.tm.ParseToken(tokenString)
	if err != nil {
		return err
	}

	user, err := h.usersRepo.GetByID(ctx, token.UserID)
	if err != nil {
		return err
	}

	if token.String != user.Token {
		return ErrTokenExpired
	}

	user.Token = ""

	if err := h.usersRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
