package usecases

import (
	"context"
	"document_manager/internal/application/domain"
	"fmt"
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

func (h *RegisterUserHandler) Execute(ctx context.Context, adminToken, login, password string) (string, error) {
	if h.adminToken != adminToken {
		return "", fmt.Errorf("provided admin token is incorrect")
	}

	user, err := domain.NewUserInput(login, password)
	if err != nil {
		return "", fmt.Errorf("validating user error: %v", err)
	}
	if err := h.usersRepo.Save(ctx, user); err != nil {
		return "", err
	}
	return user.Login, nil
}
