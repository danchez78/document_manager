package token_generator

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenGenerator struct {
	secretkey       []byte
	tokenTTLMinutes int
}

func NewTokenGenerator(cfg Config) *TokenGenerator {
	return &TokenGenerator{
		secretkey:       []byte(cfg.SecretKey),
		tokenTTLMinutes: cfg.TokenTTLMinutes,
	}
}

func (g *TokenGenerator) GenerateToken(userID string) (*Token, error) {
	tokenExpirationTime := time.Now().Add(time.Minute * time.Duration(g.tokenTTLMinutes)).Unix()

	tokenClaims := jwt.MapClaims{
		"user_id":         userID,
		"expiration_time": tokenExpirationTime,
	}

	accessToken, err := g.generateToken(tokenClaims)
	if err != nil {
		return nil, err
	}

	return &Token{
		UserID:         userID,
		String:         accessToken,
		ExpirationTime: tokenExpirationTime,
	}, nil
}

func (g *TokenGenerator) ParseToken(tokenString string) (*Token, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.secretkey), nil
	})
	if err != nil {
		return nil, err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("provided token with incorrect user id type")
	}

	expirationTime, ok := claims["expiration_time"].(float64)
	if !ok {
		return nil, errors.New("provided token with incorrect expiration type")
	}

	return &Token{
		UserID:         userID,
		String:         tokenString,
		ExpirationTime: int64(expirationTime),
	}, nil
}

func (g *TokenGenerator) generateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString(g.secretkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
