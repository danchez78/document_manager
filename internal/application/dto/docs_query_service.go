package dto

import (
	"context"
)

type DocsQueryService interface {
	GetDocsInfo(ctx context.Context, login string, limit int, docFilters *DocFilters) ([]*DocInfo, error)
	GetDocByID(ctx context.Context, id string) (*Doc, error)
}
