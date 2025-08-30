package usecases

import "errors"

var (
	ErrIncorrectPasswordProvided = errors.New("provided incorrect password")
	ErrTokenExpired              = errors.New("provided token expired")
	ErrEmptyToken                = errors.New("provided token is empty")
	ErrNoAccessToDoc             = errors.New("provided user has no access to requested document")
	ErrNoAccessToDeleteDoc       = errors.New("provided user has no access to delete requested document")
)
