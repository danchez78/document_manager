package domain

import "context"

type UsersRepository interface {
	Save(ctx context.Context, user *UserInput) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByLogin(ctx context.Context, login string) (*User, error)
	Update(ctx context.Context, user *User) error
}
