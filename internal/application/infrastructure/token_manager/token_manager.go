package token_manager

import (
	"document_manager/internal/application/domain"
	"document_manager/internal/common/token_generator"
)

type tokenManager struct {
	tg *token_generator.TokenGenerator
}

var (
	_ domain.TokenManager = (*tokenManager)(nil)
)

func NewTokenManager(tg *token_generator.TokenGenerator) *tokenManager {
	if tg == nil {
		panic("token generator is nil")
	}

	return &tokenManager{tg: tg}
}

func (tm *tokenManager) GenerateToken(userID string) (string, error) {
	accessToken, err := tm.tg.GenerateToken(userID)
	if err != nil {
		return "", err
	}

	return accessToken.String, nil
}

func (tm *tokenManager) ParseToken(tokenString string) (*domain.Token, error) {
	token, err := tm.tg.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return &domain.Token{
		UserID:         token.UserID,
		String:         token.String,
		ExpirationTime: token.ExpirationTime,
	}, nil
}
