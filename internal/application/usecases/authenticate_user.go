package usecases

import (
	"context"

	"document_manager/internal/application/domain"
)

type AuthenticateUserHandler struct {
	usersRepo domain.UsersRepository
	tm        domain.TokenManager
}

func NewAuthenticateUserHandler(usersRepo domain.UsersRepository, tm domain.TokenManager) *AuthenticateUserHandler {
	if usersRepo == nil {
		panic("users repository is nil")
	}
	if tm == nil {
		panic("token manager is nil")
	}

	return &AuthenticateUserHandler{usersRepo: usersRepo, tm: tm}
}

func (h *AuthenticateUserHandler) Execute(ctx context.Context, userInput *domain.UserInput) (string, error) {
	user, err := h.usersRepo.GetByLogin(ctx, userInput.Login)
	if err != nil {
		return "", err
	}

	if err := userInput.ComparePasswordAndHash(user.HashedPassword); err != nil {
		return "", err
	}

	token, err := h.tm.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	user.Authenticate(token)
	if err := h.usersRepo.Update(ctx, user); err != nil {
		return "", err
	}

	return token, nil
}
