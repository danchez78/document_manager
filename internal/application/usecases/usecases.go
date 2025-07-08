package usecases

import (
	"document_manager/internal/application/domain"
	"document_manager/internal/application/dto"
)

type Users struct {
	RegisterUserHandler *RegisterUserHandler
	AuthUserHandler     *AuthUserHandler
	DeauthUserHandler   *DeauthUserHandler
}

type Docs struct {
	UploadDocHandler   *UploadDocHandler
	GetDocsInfoHandler *GetDocsInfoHandler
	GetDocHandler      *GetDocHandler
	DeleteDocHandler   *DeleteDocHandler
}

type UseCases struct {
	Users *Users
	Docs  *Docs
}

func NewUseCases(
	adminToken string,
	usersRepo domain.UsersRepository,
	docsRepo domain.DocsRepository,
	docsQueryService dto.DocsQueryService,
	docsCache domain.DocsCache,
	tm domain.TokenManager,
) *UseCases {
	return &UseCases{
		Users: &Users{
			RegisterUserHandler: NewRegisterUserHandler(adminToken, usersRepo),
			AuthUserHandler:     NewAuthUserHandler(usersRepo, tm),
			DeauthUserHandler:   NewDeauthUserHandler(usersRepo, tm),
		},
		Docs: &Docs{
			UploadDocHandler:   NewUploadDocHandler(usersRepo, docsRepo, docsCache, tm),
			GetDocsInfoHandler: NewGetDocsInfoHandler(usersRepo, docsQueryService, tm),
			GetDocHandler:      NewGetDocHandler(usersRepo, docsRepo, docsQueryService, docsCache, tm),
			DeleteDocHandler:   NewDeleteDocHandler(usersRepo, docsRepo, tm),
		},
	}
}
