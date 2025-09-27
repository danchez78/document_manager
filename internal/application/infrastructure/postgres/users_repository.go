package postgres

import (
	"context"
	"errors"
	"fmt"

	"document_manager/internal/application/domain"
	"document_manager/internal/application/infrastructure/postgres/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type usersRepository struct {
	client *Client
}

var (
	_ domain.UsersRepository = (*usersRepository)(nil)
)

func NewUsersRepository(client *Client) domain.UsersRepository {
	if client == nil {
		panic("postgres client is nil")
	}

	return &usersRepository{client: client}
}

func (r *usersRepository) Save(ctx context.Context, user *domain.UserInput) error {
	mUser, err := models.NewUserFromDomainInput(user)
	if err != nil {
		return err
	}

	q := "INSERT INTO document_manager.users (id, login, hashed_password, token) VALUES ($1, $2, $3, $4)"
	if _, err := r.client.Exec(ctx, q, mUser.ID, mUser.Login, mUser.HashedPassword, mUser.Token); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return ErrDuplicateUser
			}
		}
		return fmt.Errorf("inserting user got error: %v", err)
	}

	return nil
}

func (r *usersRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var mUser models.User

	q := `
	SELECT id, login, hashed_password, token
	FROM document_manager.users
	WHERE id = $1
	`
	if err := r.client.QueryRow(ctx, q, id).Scan(&mUser.ID, &mUser.Login, &mUser.HashedPassword, &mUser.Token); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoProvidedUser
		} else {
			return nil, fmt.Errorf("getting user by id got error: %v", err)
		}
	}
	return mUser.ToDomain(), nil
}

func (r *usersRepository) GetByLogin(ctx context.Context, login string) (*domain.User, error) {
	var mUser models.User

	q := `
	SELECT id, login, hashed_password, token
	FROM document_manager.users
	WHERE login = $1
	`
	if err := r.client.QueryRow(ctx, q, login).Scan(&mUser.ID, &mUser.Login, &mUser.HashedPassword, &mUser.Token); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoProvidedUser
		} else {
			return nil, fmt.Errorf("getting user by id got error: %v", err)
		}
	}
	return mUser.ToDomain(), nil
}

func (r *usersRepository) Update(ctx context.Context, user *domain.User) error {
	mUser := models.NewUserFromDomain(user)

	q := `
	UPDATE document_manager.users
	SET token = $2
	WHERE id = $1
	`
	if _, err := r.client.Exec(ctx, q, mUser.ID, mUser.Token); err != nil {
		return fmt.Errorf("updating user got error: %v", err)
	}

	return nil
}
