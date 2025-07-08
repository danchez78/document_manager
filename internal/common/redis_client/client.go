package redis_client

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	*redis.Client
}

type Config struct {
	ConnectionUrl string `yaml:"connection_url"`
}

func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	opt, err := redis.ParseURL(cfg.ConnectionUrl)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	if err := client.Set(ctx, "foo", "bar", 0).Err(); err != nil {
		return nil, err
	}

	if _, err := client.Get(ctx, "foo").Result(); err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
