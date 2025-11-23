package domain

import "errors"

var (
	ErrEmptyToken                = errors.New("provided token is empty")
	ErrTokenExpired              = errors.New("provided token expired")
	ErrInvalidHash               = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion       = errors.New("incompatible version of argon2")
	ErrIncorrectPasswordProvided = errors.New("provided incorrect password")
	ErrNoAccessToDoc             = errors.New("provided user has no access to requested document")
	ErrNoAccessToDeleteDoc       = errors.New("provided user has no access to delete requested document")
	ErrIncorrectAdminToken       = errors.New("provided admin token is incorrect")
)
