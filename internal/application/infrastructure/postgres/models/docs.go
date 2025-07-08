package models

import (
	"document_manager/internal/application/domain"
	"document_manager/internal/application/dto"
	"time"

	"github.com/google/uuid"
)

type DocInput struct {
	ID      string
	Data    []byte
	Name    string
	File    bool
	Public  bool
	Created string
	Mime    string
	Grant   []string
}

func NewDocInputFromDomain(doc *domain.DocInput) *DocInput {
	created := time.Now().Format("2006-01-02 15:04:05")
	return &DocInput{
		ID:      uuid.New().String(),
		Data:    doc.Data,
		Name:    doc.Name,
		File:    doc.File,
		Public:  doc.Public,
		Created: created,
		Mime:    doc.Mime,
		Grant:   doc.Grant,
	}
}

type DocInfo struct {
	ID      string
	Name    string
	Mime    string
	File    bool
	Public  bool
	Created string
	Grant   []string
}

func (i *DocInfo) ToDomain() *domain.DocInfo {
	return &domain.DocInfo{
		ID:     i.ID,
		Name:   i.Name,
		File:   i.File,
		Public: i.Public,
		Mime:   i.Mime,
		Grant:  i.Grant,
	}
}

type DocInfoPreview struct {
	ID      string
	Name    string
	Mime    string
	File    bool
	Public  bool
	Created string
	Grant   []string
}

func (m *DocInfoPreview) ToDTO() *dto.DocInfo {
	return &dto.DocInfo{
		ID:      m.ID,
		Name:    m.Name,
		Mime:    m.Mime,
		File:    m.File,
		Public:  m.Public,
		Created: m.Created,
		Grant:   m.Grant,
	}
}

type DocInfoPreviews []*DocInfoPreview

func (ms DocInfoPreviews) ToDTO() []*dto.DocInfo {
	docsInfo := make([]*dto.DocInfo, 0, len(ms))
	for _, m := range ms {
		docsInfo = append(docsInfo, m.ToDTO())
	}
	return docsInfo
}

type DocPreview struct {
	Data []byte
	Mime string
}

func (p *DocPreview) ToDTO() *dto.Doc {
	return &dto.Doc{
		Data: p.Data,
		Mime: p.Mime,
	}
}
