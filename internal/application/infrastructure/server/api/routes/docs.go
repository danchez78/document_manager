package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"document_manager/internal/application/infrastructure/server"
	"document_manager/internal/application/infrastructure/server/api/controllers"
	"document_manager/internal/application/infrastructure/server/api/error_handlers"
	"document_manager/internal/application/infrastructure/server/api/views"
	"document_manager/internal/application/usecases"
)

type docsHandler struct {
	uc *usecases.UseCases
}

// uploadDoc godoc
//
//	@Summary		Upload doc
//	@Description	Uploads document
//	@Tags			docs
//	@Accept			json
//	@Produce		json
//	@Param			file	formData	file	true	"File of document"
//	@Param			meta	formData	file	true	"Meta data of document"
//	@Success		200		{object}	views.Response[views.UploadDocResponse]
//	@Router			/docs [post]
func (h *docsHandler) uploadDoc(c echo.Context) error {
	ctx := c.Request().Context()

	req, err := controllers.BindUploadDocRequest(c)
	if err != nil {
		log.Printf("failed to decode request. Reason: %v", err)

		return views.ReturnError(c, http.StatusBadRequest, err)
	}

	docDomain := req.UploadDocController.ToDomainDoc()
	err = h.uc.Docs.UploadDocHandler.Execute(ctx, docDomain, req.UploadDocController.Meta.Token)
	if err != nil {
		log.Printf("failed to upload doc. Reason: %v", err)

		return error_handlers.HandleError(c, err)
	}

	return views.ReturnResponse(c, views.NewUploadDocResponse(docDomain.Name))
}

// getDocsInfo godoc
//
//	@Summary		Get docs info
//	@Description	Returns documents information
//	@Tags			docs
//	@Accept			json
//	@Produce		json
//	@Param			token	query		string	true	"User token"
//	@Param			login	query		string	false	"User login"
//	@Param			limit	query		integer	false	"Count of documents to return. Default: 10"
//	@Param			name	query		string	false	"Name of document"
//	@Param			mime	query		string	false	"Mime of document"
//	@Param			file	query		bool	false	"Is file"
//	@Param			public	query		bool	false	"Is public"
//	@Param			created	query		string	false	"Creation time of document"
//	@Success		200		{object}	views.Response[views.GetDocsInfoResponse]
//	@Router			/docs [get]
func (h *docsHandler) getDocsInfo(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.GetDocsInfoRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("failed to decode request. Reason: %v", err)

		return views.ReturnError(c, http.StatusBadRequest, err)
	}

	docsInfo, err := h.uc.Docs.GetDocsInfoHandler.Execute(ctx, req.Token, req.Login, req.Limit, req.DocFilters.ToDomain())
	if err != nil {
		log.Printf("failed to get docs info. Reason: %v", err)

		return error_handlers.HandleError(c, err)
	}

	return views.ReturnResponse(c, views.NewGetDocsInfoResponse(docsInfo))
}

// getDoc godoc
//
//	@Summary		Get doc
//	@Description	Returns document
//	@Tags			docs
//	@Accept			json
//	@Produce		json
//	@Param			token	query	string	true	"Token of user"
//	@Param			id		path	string	true	"ID of document"
//	@Success		200		{file}	file
//	@Router			/docs/{id} [get]
func (h *docsHandler) getDoc(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.GetDocRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("failed to decode request. Reason: %v", err)

		return views.ReturnError(c, http.StatusBadRequest, err)
	}

	doc, err := h.uc.Docs.GetDocHandler.Execute(ctx, req.Token, req.ID)
	if err != nil {
		log.Printf("failed to get doc. Reason: %v", err)

		return error_handlers.HandleError(c, err)
	}

	if doc.Mime == "application/json" {
		var data map[string]interface{}

		if err := json.Unmarshal(doc.Data, &data); err != nil {
			log.Printf("failed to unmarshal document. Reason: %v", err)

			return views.ReturnError(c, http.StatusInternalServerError, err)
		}

		return views.ReturnData(c, data)
	}

	return views.ReturnFile(c, doc.Mime, doc.Data)
}

// deleteDoc godoc
//
//	@Summary		Delete doc
//	@Description	Deletes document
//	@Tags			docs
//	@Accept			json
//	@Produce		json
//	@Param			token	query	string	true	"Token of user"
//	@Param			id		path	string	true	"ID of document"
//	@Success		200		{file}	file
//	@Router			/docs/{id} [delete]
func (h *docsHandler) deleteDoc(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.DeleteDocRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("failed to decode request. Reason: %v", err)

		return views.ReturnError(c, http.StatusBadRequest, err)
	}

	err := h.uc.Docs.DeleteDocHandler.Execute(ctx, req.Token, req.ID)
	if err != nil {
		log.Printf("failed to delete doc. Reason: %v", err)

		return views.ReturnError(c, http.StatusInternalServerError, err)
	}

	return views.ReturnResponse(c, views.NewDeleteDocResponse(req.ID))
}

func makeDocsRoutes(srv *server.Server, uc *usecases.UseCases) {
	sg := srv.BasePath()
	h := docsHandler{uc: uc}

	{
		sg := sg.Group("/docs")

		sg.POST("", h.uploadDoc)
		sg.GET("", h.getDocsInfo)
		sg.GET("/:id", h.getDoc)
		sg.DELETE("/:id", h.deleteDoc)
	}
}
