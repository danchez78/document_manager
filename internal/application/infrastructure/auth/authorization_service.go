package auth

import (
	"context"
	"slices"

	"document_manager/internal/application/domain"
)

type authorizationService struct {
	docsRepo domain.DocsRepository
}

var (
	_ domain.AuthorizationService = (*authorizationService)(nil)
)

func NewAuthorizationService(docsRepo domain.DocsRepository) domain.AuthorizationService {
	if docsRepo == nil {
		panic("docs repository is nil")
	}
	return &authorizationService{
		docsRepo: docsRepo,
	}
}

func (s *authorizationService) CheckDocAccess(ctx context.Context, user *domain.User, docID string) error {
	docInfo, err := s.docsRepo.GetDocInfoByID(ctx, docID)
	if err != nil {
		return err
	}

	if !(docInfo.Public || slices.Contains(docInfo.Grant, user.Login)) {
		return domain.ErrNoAccessToDoc
	}

	return nil
}

func (s *authorizationService) CheckDocDeleteAccess(ctx context.Context, user *domain.User, docID string) error {
	docInfo, err := s.docsRepo.GetDocInfoByID(ctx, docID)
	if err != nil {
		return err
	}

	if !slices.Contains(docInfo.Grant, user.Login) {
		return domain.ErrNoAccessToDeleteDoc
	}

	return nil
}
