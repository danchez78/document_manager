package usecases

import (
	"context"
	"document_manager/internal/application/domain"
)

type GetDocsInfoHandler struct {
	authenticationService domain.AuthenticationService
	docsRepo              domain.DocsRepository
}

func NewGetDocsInfoHandler(
	authenticationService domain.AuthenticationService,
	docsRepo domain.DocsRepository,
) *GetDocsInfoHandler {
	if authenticationService == nil {
		panic("authentication service is nil")
	}
	if docsRepo == nil {
		panic("docs repository is nil")
	}
	return &GetDocsInfoHandler{
		authenticationService: authenticationService,
		docsRepo:              docsRepo,
	}
}

func (h *GetDocsInfoHandler) Execute(ctx context.Context, tokenString string, limit int, docFilters *domain.DocFilters) ([]*domain.DocInfo, error) {
	user, err := h.authenticationService.ValidateToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}

	docsInfo, err := h.docsRepo.GetDocsInfo(ctx, user.Login, limit, docFilters)
	if err != nil {
		return nil, err
	}

	return docsInfo, nil
}
