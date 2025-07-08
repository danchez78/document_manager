package postgres

import "errors"

var (
	ErrNoProvidedUser = errors.New("user with provided id not found")
	ErrDuplicateUser  = errors.New("user with provided login already exists")
	ErrNoProvidedDoc  = errors.New("document with provided id not found")
)
