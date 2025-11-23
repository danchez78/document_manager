package usecases

import (
	"context"
	"document_manager/internal/application/domain"
)

type UploadDocHandler struct {
	authenticationService domain.AuthenticationService
	docsRepo              domain.DocsRepository
	docsCache             domain.DocsCache
}

func NewUploadDocHandler(
	authenticationService domain.AuthenticationService,
	docsRepo domain.DocsRepository,
	docsCache domain.DocsCache,
) *UploadDocHandler {
	if authenticationService == nil {
		panic("authentication service is nil")
	}
	if docsRepo == nil {
		panic("docs repository is nil")
	}
	if docsCache == nil {
		panic("docs cache is nil")
	}
	return &UploadDocHandler{
		authenticationService: authenticationService,
		docsRepo:              docsRepo,
		docsCache:             docsCache,
	}
}

func (h *UploadDocHandler) Execute(ctx context.Context, doc *domain.DocInput, tokenString string) error {
	if _, err := h.authenticationService.ValidateToken(ctx, tokenString); err != nil {
		return err
	}

	docID, err := h.docsRepo.Save(ctx, doc)
	if err != nil {
		return err
	}

	return h.docsCache.Set(ctx, docID, doc.Data)
}
