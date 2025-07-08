package domain

type TokenManager interface {
	GenerateToken(userID string) (string, error)
	ParseToken(tokenString string) (*Token, error)
}
