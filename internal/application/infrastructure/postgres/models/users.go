package models

import (
	"document_manager/internal/application/domain"

	"github.com/google/uuid"
)

type User struct {
	ID             string
	Login          string
	HashedPassword string
	Token          string
}

func NewUserFromDomainInput(user *domain.UserInput) (*User, error) {
	hashedPassword, err := user.HashPassword()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:             uuid.New().String(),
		Login:          user.Login,
		HashedPassword: hashedPassword,
		Token:          "",
	}, nil
}

func NewUserFromDomain(user *domain.User) *User {
	return &User{
		ID:             user.ID,
		Login:          user.Login,
		HashedPassword: user.HashedPassword,
		Token:          user.Token,
	}
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		ID:             u.ID,
		Login:          u.Login,
		HashedPassword: u.HashedPassword,
		Token:          u.Token,
	}
}
