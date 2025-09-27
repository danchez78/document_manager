package usecases

import (
	"context"

	"document_manager/internal/application/domain"
)

type AuthUserHandler struct {
	usersRepo domain.UsersRepository
}

func NewAuthUserHandler(usersRepo domain.UsersRepository) *AuthUserHandler {
	if usersRepo == nil {
		panic("users repository is nil")
	}

	return &AuthUserHandler{usersRepo: usersRepo}
}

func (h *AuthUserHandler) Execute(ctx context.Context, login, password string) (string, error) {
	userInput, err := domain.NewUserInput(login, password)
	if err != nil {
		return "", err
	}

	user, err := h.usersRepo.GetByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	match, err := userInput.СomparePasswordAndHash(user.HashedPassword)
	if err != nil {
		return "", err
	}
	if !match {
		return "", ErrIncorrectPasswordProvided
	}

	token, err := domain.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	user.Token = token.String

	if err := h.usersRepo.Update(ctx, user); err != nil {
		return "", err
	}

	return user.Token, nil
}
