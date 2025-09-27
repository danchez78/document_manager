package routes

import (
	"document_manager/internal/application/infrastructure/server"
	"document_manager/internal/application/usecases"
)

func Make(
	srv *server.Server,
	uc *usecases.UseCases,
) {
	makeUsersRoutes(srv, uc)
	makeDocsRoutes(srv, uc)
}
