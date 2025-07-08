package postgres

import (
	"context"
	"errors"
	"fmt"

	"document_manager/internal/application/dto"
	"document_manager/internal/application/infrastructure/postgres/models"
	"document_manager/internal/common/postgres_client"

	"github.com/jackc/pgx/v5"
)

type docsQueryService struct {
	client *postgres_client.Client
}

var (
	_ dto.DocsQueryService = (*docsQueryService)(nil)
)

func NewDocsQueryService(client *postgres_client.Client) dto.DocsQueryService {
	if client == nil {
		panic("postgres client is nil")
	}

	return &docsQueryService{client: client}
}

func (r *docsQueryService) GetDocsInfo(ctx context.Context, login string, limit int, docFilters *dto.DocFilters) ([]*dto.DocInfo, error) {
	nameMatch := docFilters.Name
	mimeMatch := docFilters.Mime
	createdMatch := docFilters.Created

	if nameMatch != "" {
		nameMatch = "%" + nameMatch + "%"
	}

	if mimeMatch != "" {
		mimeMatch = "%" + mimeMatch + "%"
	}

	if createdMatch != "" {
		createdMatch = "%" + createdMatch + "%"
	}

	grants := []string{login}

	if limit == 0 {
		limit = 10
	}

	q := `
		WITH filtered AS (
			SELECT d.id, d.name, d.mime, d.file, d.public, d.created, ARRAY_AGG(g.user_login) AS grants
			FROM document_manager.docs d
			JOIN document_manager.grants g
			ON d.id=g.doc_id
			WHERE ('' = $1 OR d.name ILIKE $1)
			AND ('' = $2 OR d.mime ILIKE $2)
			AND ('' = $3 OR d.created LIKE $3)
			AND ($4::bool IS NULL OR d.file = $4)
			AND ($5::bool IS NULL OR d.public = $5)
			GROUP BY d.id
		)
		SELECT * FROM filtered
		WHERE (cardinality($6::varchar[]) IS NULL OR grants @> $6::varchar[])
		ORDER BY name, created DESC
		LIMIT $7
		`

	rows, err := r.client.Query(ctx, q, nameMatch, mimeMatch, createdMatch, docFilters.File, docFilters.Public, grants, limit)
	if err != nil {
		return nil, fmt.Errorf("getting docs info got error: %v", err)
	}

	mDocsInfo := make(models.DocInfoPreviews, 0)
	for rows.Next() {
		mDocInfo := new(models.DocInfoPreview)
		if err := rows.Scan(
			&mDocInfo.ID,
			&mDocInfo.Name,
			&mDocInfo.Mime,
			&mDocInfo.File,
			&mDocInfo.Public,
			&mDocInfo.Created,
			&mDocInfo.Grant,
		); err != nil {
			return nil, fmt.Errorf("getting doc info got error: %v", err)
		}
		mDocsInfo = append(mDocsInfo, mDocInfo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("listing docs info got error: %v", err)
	}

	return mDocsInfo.ToDTO(), nil

}

func (r *docsQueryService) GetDocByID(ctx context.Context, id string) (*dto.Doc, error) {
	var mDoc models.DocPreview
	q := `
		SELECT d.data, d.mime
		FROM document_manager.docs d
		JOIN document_manager.grants g
		ON d.id=g.doc_id
		WHERE d.id = $1
		GROUP BY d.id
		`

	if err := r.client.QueryRow(ctx, q, id).Scan(&mDoc.Data, &mDoc.Mime); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoProvidedUser
		} else {
			return nil, fmt.Errorf("getting user by id got error: %v", err)
		}
	}

	return mDoc.ToDTO(), nil

}
