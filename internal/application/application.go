package application

import (
	"context"

	"document_manager/config"
	"document_manager/internal/application/infrastructure/api/routes"
	"document_manager/internal/application/infrastructure/api/server"
	"document_manager/internal/application/infrastructure/auth"
	"document_manager/internal/application/infrastructure/postgres"
	"document_manager/internal/application/infrastructure/redis"
	"document_manager/internal/application/infrastructure/token"
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

	tg := token.NewTokenGenerator(cfg.TokenGenerator)
	tm := token.NewTokenManager(tg)

	usersRepo := postgres.NewUsersRepository(pg_client)
	docRepo := postgres.NewDocsRepository(pg_client)

	docsCache := redis.NewDocsCache(rd_client)

	authenticationService := auth.NewAuthenticationService(usersRepo, tm)
	authorizationService := auth.NewAuthorizationService(docRepo)

	uc := usecases.NewUseCases(cfg.AdminToken, usersRepo, docRepo, docsCache, tm, authenticationService, authorizationService)
	routes.Make(srv, uc)

	return nil
}
