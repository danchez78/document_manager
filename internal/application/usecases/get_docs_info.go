package usecases

import (
	"context"

	"document_manager/internal/application/domain"
)

type GetDocsInfoHandler struct {
	usersRepo domain.UsersRepository
	docsRepo  domain.DocsRepository
	tm        domain.TokenManager
}

func NewGetDocsInfoHandler(
	usersRepo domain.UsersRepository,
	docsRepo domain.DocsRepository,
	tm domain.TokenManager,
) *GetDocsInfoHandler {
	if usersRepo == nil {
		panic("users repository is nil")
	}

	if docsRepo == nil {
		panic("docs repository is nil")
	}

	if tm == nil {
		panic("token manager is nil")
	}

	return &GetDocsInfoHandler{usersRepo: usersRepo, docsRepo: docsRepo, tm: tm}
}

func (h *GetDocsInfoHandler) Execute(ctx context.Context, tokenString, login string, limit int, docFilters *domain.DocFilters) ([]*domain.DocInfo, error) {
	if tokenString == "" {
		return nil, ErrEmptyToken
	}

	token, err := h.tm.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if token.IsExpired() {
		return nil, ErrTokenExpired
	}

	user, err := h.usersRepo.GetByID(ctx, token.UserID)
	if err != nil {
		return nil, err
	}

	if login == "" {
		login = user.Login
	}

	if token.String != user.Token {
		return nil, ErrTokenExpired
	}

	docsInfo, err := h.docsRepo.GetDocsInfo(ctx, login, limit, docFilters)
	if err != nil {
		return nil, err
	}

	return docsInfo, nil
}
