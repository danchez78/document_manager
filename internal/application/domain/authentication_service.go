package domain

import "context"

type AuthenticationService interface {
	ValidateToken(ctx context.Context, tokenString string) (*User, error)
}
