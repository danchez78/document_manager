package application

import (
	"context"
	"document_manager/config"
	"document_manager/internal/application/infrastructure/api/routes"
	"document_manager/internal/application/infrastructure/postgres"
	"document_manager/internal/application/infrastructure/redis"
	"document_manager/internal/application/infrastructure/token_manager"
	"document_manager/internal/application/usecases"
	"document_manager/internal/common/postgres_client"
	"document_manager/internal/common/redis_client"
	"document_manager/internal/common/server"
	"document_manager/internal/common/token_generator"
)

func Init(
	ctx context.Context,
	srv *server.Server,
	cfg config.Config,
) error {
	pg_client, err := postgres_client.NewClient(ctx, cfg.Postgres)
	if err != nil {
		return err
	}

	rd_client, err := redis_client.NewClient(ctx, cfg.RedisClient)
	if err != nil {
		return err
	}

	tg := token_generator.NewTokenGenerator(cfg.TokenGenerator)
	tm := token_manager.NewTokenManager(tg)

	usersRepo := postgres.NewUsersRepository(pg_client)
	docRepo := postgres.NewDocsRepository(pg_client)
	docsQueryService := postgres.NewDocsQueryService(pg_client)

	docsCache := redis.NewDocsCache(rd_client)

	uc := usecases.NewUseCases(cfg.AdminToken, usersRepo, docRepo, docsQueryService, docsCache, tm)
	routes.Make(srv, uc)

	return nil
}
