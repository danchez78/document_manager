package usecases

import (
	"context"

	"document_manager/internal/application/domain"
)

type UploadDocHandler struct {
	usersRepo domain.UsersRepository
	docsRepo  domain.DocsRepository
	docsCache domain.DocsCache
}

func NewUploadDocHandler(
	usersRepo domain.UsersRepository,
	docsRepo domain.DocsRepository,
	docsCache domain.DocsCache,
) *UploadDocHandler {
	if usersRepo == nil {
		panic("user repository is nil")
	}

	if docsRepo == nil {
		panic("docs repository is nil")
	}

	if docsCache == nil {
		panic("docs cache is nil")
	}

	return &UploadDocHandler{usersRepo: usersRepo,
		docsRepo:  docsRepo,
		docsCache: docsCache,
	}
}

func (h *UploadDocHandler) Execute(ctx context.Context, doc *domain.DocInput, tokenString string) error {
	if tokenString == "" {
		return ErrEmptyToken
	}

	token, err := domain.ParseToken(tokenString)
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

	docID, err := h.docsRepo.Save(ctx, doc)
	if err != nil {
		return err
	}

	if err := h.docsCache.Set(ctx, docID, doc.Data); err != nil {
		return err
	}

	return nil
}
