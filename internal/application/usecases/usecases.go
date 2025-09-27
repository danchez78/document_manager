package usecases

import (
	"document_manager/internal/application/domain"
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
	docsCache domain.DocsCache,
) *UseCases {
	return &UseCases{
		Users: &Users{
			RegisterUserHandler: NewRegisterUserHandler(adminToken, usersRepo),
			AuthUserHandler:     NewAuthUserHandler(usersRepo),
			DeauthUserHandler:   NewDeauthUserHandler(usersRepo),
		},
		Docs: &Docs{
			UploadDocHandler:   NewUploadDocHandler(usersRepo, docsRepo, docsCache),
			GetDocsInfoHandler: NewGetDocsInfoHandler(usersRepo, docsRepo),
			GetDocHandler:      NewGetDocHandler(usersRepo, docsRepo, docsCache),
			DeleteDocHandler:   NewDeleteDocHandler(usersRepo, docsRepo),
		},
	}
}
