package auth

import (
	"context"
	"document_manager/internal/application/domain"
)

type authenticationService struct {
	usersRepo domain.UsersRepository
	tm        domain.TokenManager
}

var (
	_ domain.AuthenticationService = (*authenticationService)(nil)
)

func NewAuthenticationService(usersRepo domain.UsersRepository, tm domain.TokenManager) domain.AuthenticationService {
	if usersRepo == nil {
		panic("users repository is nil")
	}
	if tm == nil {
		panic("token manager is nil")
	}
	return &authenticationService{
		usersRepo: usersRepo,
		tm:        tm,
	}
}

func (s *authenticationService) ValidateToken(ctx context.Context, tokenString string) (*domain.User, error) {
	if tokenString == "" {
		return nil, domain.ErrEmptyToken
	}

	token, err := s.tm.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if token.IsExpired() {
		return nil, domain.ErrTokenExpired
	}

	user, err := s.usersRepo.GetByID(ctx, token.UserID)
	if err != nil {
		return nil, err
	}

	if token.String != user.Token {
		return nil, domain.ErrTokenExpired
	}

	return user, nil
}
