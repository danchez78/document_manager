package redis

import (
	"context"
	"document_manager/internal/application/domain"
	"document_manager/internal/common/redis_client"

	"github.com/redis/go-redis/v9"
)

type docsCache struct {
	client *redis_client.Client
}

var (
	_ domain.DocsCache = (*docsCache)(nil)
)

func NewDocsCache(client *redis_client.Client) domain.DocsCache {
	if client == nil {
		panic("redis client is nil")
	}

	return &docsCache{client: client}
}

func (c *docsCache) Set(ctx context.Context, id string, data []byte) error {
	if err := c.client.Set(ctx, id, data, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (c *docsCache) Get(ctx context.Context, id string) ([]byte, error) {
	data, err := c.client.Get(ctx, id).Bytes()
	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return data, nil
}
