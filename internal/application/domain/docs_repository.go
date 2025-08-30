package domain

import "context"

type DocsRepository interface {
	Save(ctx context.Context, doc *DocInput) (string, error)
	GetDocInfoByID(ctx context.Context, id string) (*DocInfo, error)
	GetDocsInfo(ctx context.Context, login string, limit int, docFilters *DocFilters) ([]*DocInfo, error)
	GetDocByID(ctx context.Context, id string) (*Doc, error)
	Delete(ctx context.Context, id string) error
}
