package views

import "document_manager/internal/application/dto"

type UploadDocResponse struct {
	File string `json:"file"`
}

func NewUploadDocResponse(fileName string) UploadDocResponse {
	return UploadDocResponse{File: fileName}
}

type DocInfo struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Mime    string   `json:"mime"`
	File    bool     `json:"file"`
	Public  bool     `json:"public"`
	Created string   `json:"created"`
	Grant   []string `json:"grant"`
}

type GetDocsInfoResponse struct {
	Docs []DocInfo `json:"docs"`
}

func NewGetDocsInfoResponse(docsInfo []*dto.DocInfo) GetDocsInfoResponse {
	docs := make([]DocInfo, 0, len(docsInfo))
	for _, docInfo := range docsInfo {
		docs = append(docs, DocInfo(*docInfo))
	}

	return GetDocsInfoResponse{Docs: docs}
}

type DeleteDocResponse map[string]bool

func NewDeleteDocResponse(id string) DeleteDocResponse {
	return DeleteDocResponse{
		id: true,
	}
}
