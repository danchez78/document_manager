package usecases

import (
	"context"
	"slices"

	"document_manager/internal/application/domain"
)

type DeleteDocHandler struct {
	usersRepo domain.UsersRepository
	docsRepo  domain.DocsRepository
	tm        domain.TokenManager
}

func NewDeleteDocHandler(
	usersRepo domain.UsersRepository,
	docsRepo domain.DocsRepository,
	tm domain.TokenManager,
) *DeleteDocHandler {
	if usersRepo == nil {
		panic("users repository is nil")
	}

	if docsRepo == nil {
		panic("docs repository is nil")
	}

	if tm == nil {
		panic("token manager is nil")
	}

	return &DeleteDocHandler{usersRepo: usersRepo, docsRepo: docsRepo, tm: tm}
}

func (h *DeleteDocHandler) Execute(ctx context.Context, tokenString, id string) error {
	if tokenString == "" {
		return ErrEmptyToken
	}

	token, err := h.tm.ParseToken(tokenString)
	if err != nil {
		return err
	}

	if token.IsExpired() {
		return ErrTokenExpired
	}

	user, err := h.usersRepo.GetByID(ctx, token.UserID)
	if err != nil {
		return err
	}

	if token.String != user.Token {
		return ErrTokenExpired
	}

	docInfo, err := h.docsRepo.Get(ctx, id)
	if err != nil {
		return err
	}

	if !(slices.Contains(docInfo.Grant, user.Login)) {
		return ErrNoAccessToDeleteDoc
	}

	if err := h.docsRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
