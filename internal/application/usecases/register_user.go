package usecases

import (
	"context"
	"document_manager/internal/application/domain"
)

type RegisterUserHandler struct {
	adminToken string
	usersRepo  domain.UsersRepository
}

func NewRegisterUserHandler(adminToken string, usersRepo domain.UsersRepository) *RegisterUserHandler {
	if usersRepo == nil {
		panic("users repository is nil")
	}
	return &RegisterUserHandler{adminToken: adminToken, usersRepo: usersRepo}
}

func (h *RegisterUserHandler) Execute(ctx context.Context, adminToken string, userInput *domain.UserInput) (string, error) {
	if h.adminToken != adminToken {
		return "", domain.ErrIncorrectAdminToken
	}

	if err := h.usersRepo.Save(ctx, userInput); err != nil {
		return "", err
	}

	return userInput.Login, nil
}
