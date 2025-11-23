package domain

import "context"

type AuthorizationService interface {
	CheckDocAccess(ctx context.Context, user *User, docID string) error
	CheckDocDeleteAccess(ctx context.Context, user *User, docID string) error
}
