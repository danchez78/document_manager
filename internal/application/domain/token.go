package domain

import "time"

type Token struct {
	UserID         string
	String         string
	ExpirationTime int64
}

func (t *Token) IsExpired() bool {
	return t.ExpirationTime < time.Now().Unix()
}
