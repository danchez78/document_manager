package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"document_manager/internal/application/infrastructure/api/controllers"
	"document_manager/internal/application/infrastructure/api/error_handlers"
	"document_manager/internal/application/infrastructure/api/views"
	"document_manager/internal/application/usecases"
	"document_manager/internal/common/server"
)

type userHandler struct {
	uc *usecases.UseCases
}

// registerUser godoc
//
//	@Summary		Register user
//	@Description	Registers user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			register_user_data	body		controllers.RegisterUserController	true	"Admin token and users login and password to register it with"
//	@Success		200					{object}	server.Response[views.RegisterUserResponse]
//	@Router			/register [post]
func (h *userHandler) registerUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.RegisterUserRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("failed to decode request. Reason: %v", err)

		return server.ReturnError(c, http.StatusBadRequest, err)
	}

	login, err := h.uc.Users.RegisterUserHandler.Execute(ctx, req.Token, req.Login, req.Password)
	if err != nil {
		log.Printf("failed to register user. Reason: %v", err)

		return error_handlers.HandleError(c, err)
	}

	return server.ReturnResponse(c, views.NewRegisterUserResponse(login))
}

// authUser godoc
//
//	@Summary		Auth user
//	@Description	Authenticates user by login and password and returns token
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			auth_user_data	body		controllers.AuthUserController	true	"Login and password"
//	@Success		200				{object}	server.Response[views.AuthUserResponse]
//	@Router			/auth [post]
func (h *userHandler) authUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.AuthUserRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("failed to decode request. Reason: %v", err)

		return server.ReturnError(c, http.StatusBadRequest, err)
	}

	token, err := h.uc.Users.AuthUserHandler.Execute(ctx, req.Login, req.Password)
	if err != nil {
		log.Printf("failed to auth user. Reason: %v", err)

		return error_handlers.HandleError(c, err)
	}

	return server.ReturnResponse(c, views.NewAuthUserResponse(token))
}

// deauthUser godoc
//
//	@Summary		Deauth user
//	@Description	Deauthenticates user by token
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	server.Response[views.DeauthUserResponse]
//	@Router			/auth/{token} [delete]
func (h *userHandler) deauthUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.DeauthUserRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("failed to decode request. Reason: %v", err)

		return server.ReturnError(c, http.StatusBadRequest, err)
	}

	err := h.uc.Users.DeauthUserHandler.Execute(ctx, req.Token)
	if err != nil {
		log.Printf("failed to deauth user. Reason: %v", err)

		return error_handlers.HandleError(c, err)
	}

	return server.ReturnResponse(c, views.NewDeauthUserResponse(req.Token))
}

func makeUsersRoutes(srv *server.Server, uc *usecases.UseCases) {
	sg := srv.BasePath()
	h := userHandler{uc: uc}

	{
		sg.POST("/register", h.registerUser)

		sg.POST("/auth", h.authUser)
		sg.DELETE("/auth/:token", h.deauthUser)
	}
}
