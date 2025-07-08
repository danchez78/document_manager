package domain

import "context"

type DocsRepository interface {
	Save(ctx context.Context, doc *DocInput) (string, error)
	Get(ctx context.Context, id string) (*DocInfo, error)
	Delete(ctx context.Context, id string) error
}
