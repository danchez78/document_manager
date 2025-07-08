package usecases

import (
	"context"
	"slices"

	"document_manager/internal/application/domain"
	"document_manager/internal/application/dto"
)

type GetDocHandler struct {
	usersRepo        domain.UsersRepository
	docsRepo         domain.DocsRepository
	docsQueryService dto.DocsQueryService
	docsCache        domain.DocsCache
	tm               domain.TokenManager
}

func NewGetDocHandler(
	usersRepo domain.UsersRepository,
	docsRepo domain.DocsRepository,
	docsQueryService dto.DocsQueryService,
	docsCache domain.DocsCache,
	tm domain.TokenManager,
) *GetDocHandler {
	if usersRepo == nil {
		panic("users repository is nil")
	}

	if docsRepo == nil {
		panic("docs repository is nil")
	}

	if docsQueryService == nil {
		panic("docs query service is nil")
	}

	if docsCache == nil {
		panic("docs cache is nil")
	}

	if tm == nil {
		panic("token manager is nil")
	}

	return &GetDocHandler{
		usersRepo:        usersRepo,
		docsRepo:         docsRepo,
		docsQueryService: docsQueryService,
		docsCache:        docsCache,
		tm:               tm,
	}
}

func (h *GetDocHandler) Execute(ctx context.Context, tokenString, id string) (*dto.Doc, error) {
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

	if token.String != user.Token {
		return nil, ErrTokenExpired
	}

	docInfo, err := h.docsRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if !(docInfo.Public || slices.Contains(docInfo.Grant, user.Login)) {
		return nil, ErrNoAccessToDoc
	}

	data, err := h.docsCache.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if data != nil {
		return &dto.Doc{Mime: docInfo.Mime, Data: data}, nil
	}

	doc, err := h.docsQueryService.GetDocByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
