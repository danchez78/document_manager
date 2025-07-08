package controllers

import (
	"document_manager/internal/application/domain"
	"document_manager/internal/application/dto"
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
)

type Meta struct {
	Name   string   `json:"name"`
	File   bool     `json:"file"`
	Public bool     `json:"public"`
	Token  string   `json:"token"`
	Mime   string   `json:"mime"`
	Grant  []string `json:"grant"`
}

type UploadDocController struct {
	File []byte
	Meta Meta
}

func (c *UploadDocController) ToDomainDoc() *domain.DocInput {
	return &domain.DocInput{
		Data:   c.File,
		Name:   c.Meta.Name,
		File:   c.Meta.File,
		Public: c.Meta.Public,
		Mime:   c.Meta.Mime,
		Grant:  c.Meta.Grant,
	}
}

type UploadDocRequest struct {
	UploadDocController
}

func BindUploadDocRequest(c echo.Context) (*UploadDocRequest, error) {
	fileData, err := readFormFile(c, "file")
	if err != nil {
		return nil, err
	}

	metaData, err := readFormFile(c, "meta")
	if err != nil {
		return nil, err
	}

	var meta Meta
	if err := json.Unmarshal(metaData, &meta); err != nil {
		return nil, err
	}

	return &UploadDocRequest{
		UploadDocController{
			File: fileData,
			Meta: meta,
		},
	}, nil
}

func readFormFile(c echo.Context, name string) ([]byte, error) {
	header, err := c.FormFile(name)
	if err != nil {
		return nil, err
	}

	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type DocFilters struct {
	Name    string `query:"name"`
	Mime    string `query:"mime"`
	File    *bool  `query:"file"`
	Public  *bool  `query:"public"`
	Created string `query:"created"`
}

func (f *DocFilters) ToDomain() *dto.DocFilters {
	return &dto.DocFilters{
		Name:    f.Name,
		Mime:    f.Mime,
		File:    f.File,
		Public:  f.Public,
		Created: f.Created,
	}
}

type GetDocsInfoRequest struct {
	Token string `query:"token"`
	Login string `query:"login"`
	Limit int    `query:"limit"`
	DocFilters
}

type GetDocRequest struct {
	ID    string `param:"id"`
	Token string `query:"token"`
}

type DeleteDocRequest struct {
	ID    string `param:"id"`
	Token string `query:"token"`
}
