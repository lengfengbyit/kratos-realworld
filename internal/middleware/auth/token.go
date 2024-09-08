package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewTokenClaims(userId int64, duration time.Duration) *TokenClaims {
	return &TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
}

// GenerateToken generates a jwt token from claims
func GenerateToken(tokenClaims *TokenClaims, secret string) string {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims).
		SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(bearerFormat, token)
}
