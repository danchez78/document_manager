package domain

import "context"

type DocsCache interface {
	Set(ctx context.Context, id string, data []byte) error
	Get(ctx context.Context, id string) ([]byte, error)
}
