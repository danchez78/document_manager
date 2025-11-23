package usecases

import (
	"document_manager/internal/application/domain"
)

type Users struct {
	RegisterUserHandler       *RegisterUserHandler
	AuthenticateUserHandler   *AuthenticateUserHandler
	DeauthenticateUserHandler *DeauthenticateUserHandler
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
	tm domain.TokenManager,
	authenticationService domain.AuthenticationService,
	authorizationService domain.AuthorizationService,
) *UseCases {
	return &UseCases{
		Users: &Users{
			RegisterUserHandler:       NewRegisterUserHandler(adminToken, usersRepo),
			AuthenticateUserHandler:   NewAuthenticateUserHandler(usersRepo, tm),
			DeauthenticateUserHandler: NewDeauthenticateUserHandler(authenticationService, usersRepo),
		},
		Docs: &Docs{
			UploadDocHandler:   NewUploadDocHandler(authenticationService, docsRepo, docsCache),
			GetDocsInfoHandler: NewGetDocsInfoHandler(authenticationService, docsRepo),
			GetDocHandler:      NewGetDocHandler(authenticationService, authorizationService, docsRepo, docsCache),
			DeleteDocHandler:   NewDeleteDocHandler(authenticationService, authorizationService, docsRepo),
		},
	}
}
