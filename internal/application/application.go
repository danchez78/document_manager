package application

import (
	"context"

	"document_manager/config"
	"document_manager/internal/application/infrastructure/postgres"
	"document_manager/internal/application/infrastructure/redis"
	"document_manager/internal/application/infrastructure/server"
	"document_manager/internal/application/infrastructure/server/api/routes"
	"document_manager/internal/application/usecases"
)

func Init(
	ctx context.Context,
	srv *server.Server,
	cfg config.Config,
) error {
	pg_client, err := postgres.NewClient(ctx, cfg.Postgres)
	if err != nil {
		return err
	}

	rd_client, err := redis.NewClient(ctx, cfg.RedisClient)
	if err != nil {
		return err
	}

	usersRepo := postgres.NewUsersRepository(pg_client)
	docRepo := postgres.NewDocsRepository(pg_client)

	docsCache := redis.NewDocsCache(rd_client)

	uc := usecases.NewUseCases(cfg.AdminToken, usersRepo, docRepo, docsCache)
	routes.Make(srv, uc)

	return nil
}
