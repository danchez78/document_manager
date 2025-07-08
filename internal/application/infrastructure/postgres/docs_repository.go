package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"

	"document_manager/internal/application/domain"
	"document_manager/internal/application/infrastructure/postgres/models"
	"document_manager/internal/common/postgres_client"

	"github.com/jackc/pgx/v5"
)

type docsRepository struct {
	client *postgres_client.Client
}

var (
	_ domain.DocsRepository = (*docsRepository)(nil)
)

func NewDocsRepository(client *postgres_client.Client) domain.DocsRepository {
	if client == nil {
		panic("postgres client is nil")
	}

	return &docsRepository{client: client}
}

func (r *docsRepository) Save(ctx context.Context, doc *domain.DocInput) (string, error) {
	mDoc := models.NewDocInputFromDomain(doc)

	tx, err := r.client.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("inserting dork got error: %v", err)
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			log.Printf("Got error during transaction rollback: %v", err)
		}
	}()

	q := "INSERT INTO document_manager.docs (id, data, name, file, public, mime, created) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	if _, err := tx.Exec(ctx, q, mDoc.ID, mDoc.Data, mDoc.Name, mDoc.File, mDoc.Public, mDoc.Mime, mDoc.Created); err != nil {
		return "", fmt.Errorf("inserting user got error: %v", err)
	}

	q = "INSERT INTO document_manager.grants (doc_id, user_login) VALUES ($1, $2)"

	for _, user_login := range mDoc.Grant {
		if _, err := tx.Exec(ctx, q, mDoc.ID, user_login); err != nil {
			log.Printf("granting doc to user got error: %v", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("inserting doc got error: %v", err)
	}

	return mDoc.ID, nil
}

func (r *docsRepository) Get(ctx context.Context, id string) (*domain.DocInfo, error) {
	var mDocInfo models.DocInfo

	q := `
		SELECT d.id, d.name, d.mime, d.file, d.public, d.created, ARRAY_AGG(g.user_login) AS grants
		FROM document_manager.docs d
		JOIN document_manager.grants g
		ON d.id=g.doc_id
		WHERE d.id = $1
		GROUP BY d.id
	`
	if err := r.client.QueryRow(ctx, q, id).Scan(
		&mDocInfo.ID,
		&mDocInfo.Name,
		&mDocInfo.Mime,
		&mDocInfo.File,
		&mDocInfo.Public,
		&mDocInfo.Created,
		&mDocInfo.Grant,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoProvidedDoc
		} else {
			return nil, fmt.Errorf("getting doc info by id got error: %v", err)
		}
	}
	return mDocInfo.ToDomain(), nil
}

func (r *docsRepository) Delete(ctx context.Context, id string) error {
	q := "DELETE FROM document_manager.docs WHERE id = $1"
	if _, err := r.client.Exec(ctx, q, id); err != nil {
		return fmt.Errorf("deleting doc got error: %v", err)
	}

	return nil
}
