package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	*pgxpool.Pool
}

type Config struct {
	ConnectionUrl string `yaml:"connection_url"`
}

func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	pool, err := pgxpool.New(ctx, cfg.ConnectionUrl)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &Client{Pool: pool}, nil
}
