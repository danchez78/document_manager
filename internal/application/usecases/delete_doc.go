package usecases

import (
	"context"
	"document_manager/internal/application/domain"
)

type DeleteDocHandler struct {
	authenticationService domain.AuthenticationService
	authorizationService  domain.AuthorizationService
	docsRepo              domain.DocsRepository
}

func NewDeleteDocHandler(
	authenticationService domain.AuthenticationService,
	authorizationService domain.AuthorizationService,
	docsRepo domain.DocsRepository,
) *DeleteDocHandler {
	if authenticationService == nil {
		panic("authentication service is nil")
	}
	if authorizationService == nil {
		panic("authorization service is nil")
	}
	if docsRepo == nil {
		panic("docs repository is nil")
	}
	return &DeleteDocHandler{
		authenticationService: authenticationService,
		authorizationService:  authorizationService,
		docsRepo:              docsRepo,
	}
}

func (h *DeleteDocHandler) Execute(ctx context.Context, tokenString, id string) error {
	user, err := h.authenticationService.ValidateToken(ctx, tokenString)
	if err != nil {
		return err
	}

	if err := h.authorizationService.CheckDocDeleteAccess(ctx, user, id); err != nil {
		return err
	}

	if err := h.docsRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
