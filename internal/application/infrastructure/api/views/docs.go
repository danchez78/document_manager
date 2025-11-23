package views

import (
	"encoding/json"

	"document_manager/internal/application/domain"
)

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
	Grants  []string `json:"grants"`
}

func NewDocInfo(docInfo *domain.DocInfo) DocInfo {
	return DocInfo{
		ID:      docInfo.ID,
		Name:    docInfo.Name,
		Mime:    docInfo.Mime,
		File:    docInfo.File,
		Public:  docInfo.Public,
		Created: docInfo.Created,
		Grants:  docInfo.Grants,
	}
}

type GetDocsInfoResponse struct {
	Docs []DocInfo `json:"docs"`
}

func NewGetDocsInfoResponse(docsInfo []*domain.DocInfo) GetDocsInfoResponse {
	docs := make([]DocInfo, 0, len(docsInfo))
	for _, docInfo := range docsInfo {
		docs = append(docs, NewDocInfo(docInfo))
	}

	return GetDocsInfoResponse{Docs: docs}
}

type DeleteDocResponse map[string]bool

func NewDeleteDocResponse(id string) DeleteDocResponse {
	return DeleteDocResponse{
		id: true,
	}
}

func NewGetDoc(doc *domain.Doc) (interface{}, bool, error) {
	if doc.Mime == "application/json" {
		var data map[string]interface{}
		if err := json.Unmarshal(doc.Data, &data); err != nil {
			return nil, false, err
		}
		return data, true, nil
	}
	return doc.Data, false, nil
}
