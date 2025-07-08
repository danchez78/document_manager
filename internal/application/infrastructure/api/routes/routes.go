package routes

import (
	"document_manager/internal/application/usecases"
	"document_manager/internal/common/server"
)

func Make(
	srv *server.Server,
	uc *usecases.UseCases,
) {
	makeUsersRoutes(srv, uc)
	makeDocsRoutes(srv, uc)
}
