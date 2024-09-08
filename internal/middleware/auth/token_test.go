package auth

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken(NewTokenClaims(1, time.Hour), "token")
	t.Log(token)
}
