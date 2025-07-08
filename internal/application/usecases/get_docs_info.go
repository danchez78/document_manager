package usecases

import (
	"context"

	"document_manager/internal/application/domain"
	"document_manager/internal/application/dto"
)

type GetDocsInfoHandler struct {
	usersRepo        domain.UsersRepository
	docsQueryService dto.DocsQueryService
	tm               domain.TokenManager
}

func NewGetDocsInfoHandler(
	usersRepo domain.UsersRepository,
	docsQueryService dto.DocsQueryService,
	tm domain.TokenManager,
) *GetDocsInfoHandler {
	if usersRepo == nil {
		panic("users repository is nil")
	}

	if docsQueryService == nil {
		panic("docs query service is nil")
	}

	if tm == nil {
		panic("token manager is nil")
	}

	return &GetDocsInfoHandler{usersRepo: usersRepo, docsQueryService: docsQueryService, tm: tm}
}

func (h *GetDocsInfoHandler) Execute(ctx context.Context, tokenString, login string, limit int, docFilters *dto.DocFilters) ([]*dto.DocInfo, error) {
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

	docsInfo, err := h.docsQueryService.GetDocsInfo(ctx, login, limit, docFilters)
	if err != nil {
		return nil, err
	}

	return docsInfo, nil
}
