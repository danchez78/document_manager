package usecases

import (
	"context"
	"document_manager/internal/application/domain"
)

type GetDocHandler struct {
	authenticationService domain.AuthenticationService
	authorizationService  domain.AuthorizationService
	docsRepo              domain.DocsRepository
	docsCache             domain.DocsCache
}

func NewGetDocHandler(
	authenticationService domain.AuthenticationService,
	authorizationService domain.AuthorizationService,
	docsRepo domain.DocsRepository,
	docsCache domain.DocsCache,
) *GetDocHandler {
	if authenticationService == nil {
		panic("authentication service is nil")
	}
	if authorizationService == nil {
		panic("authorization service is nil")
	}
	if docsRepo == nil {
		panic("docs repository is nil")
	}
	if docsCache == nil {
		panic("docs cache is nil")
	}
	return &GetDocHandler{
		authenticationService: authenticationService,
		authorizationService:  authorizationService,
		docsRepo:              docsRepo,
		docsCache:             docsCache,
	}
}

func (h *GetDocHandler) Execute(ctx context.Context, tokenString, id string) (*domain.Doc, error) {
	user, err := h.authenticationService.ValidateToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}

	if err := h.authorizationService.CheckDocAccess(ctx, user, id); err != nil {
		return nil, err
	}

	data, err := h.docsCache.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if data != nil {
		docInfo, err := h.docsRepo.GetDocInfoByID(ctx, id)
		if err != nil {
			return nil, err
		}
		return domain.NewDoc(docInfo.Mime, data), nil
	}

	return h.docsRepo.GetDocByID(ctx, id)
}
