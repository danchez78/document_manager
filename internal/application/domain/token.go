package domain

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey string
var ttlMinutes int

func SetTokenParams(tokenSecretKey string, tokenTTLMinutes int) {
	secretKey = tokenSecretKey
	ttlMinutes = tokenTTLMinutes
}

type Token struct {
	UserID         string
	String         string
	ExpirationTime int64
}

func (t *Token) IsExpired() bool {
	return t.ExpirationTime < time.Now().Unix()
}

func GenerateToken(userID string) (*Token, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(ttlMinutes)).Unix()

	tokenClaims := jwt.MapClaims{
		"user_id":         userID,
		"expiration_time": expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenClaims)

	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &Token{
		UserID:         userID,
		String:         accessToken,
		ExpirationTime: expirationTime,
	}, nil
}

func ParseToken(accessToken string) (*Token, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("provided token with incorrect user id type")
	}

	expirationTime, ok := claims["expiration_time"].(int64)
	if !ok {
		return nil, errors.New("provided token with incorrect expiration type")
	}

	return &Token{
		UserID:         userID,
		String:         accessToken,
		ExpirationTime: expirationTime,
	}, nil
}
