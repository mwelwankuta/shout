package utils

import (
	"os"

	"github.com/hako/branca"
	"github.com/joho/godotenv"
)

// store as env var

func GenerateToken(userId uint, expiry int64) (string, error) {
	godotenv.Load()

	var secret = os.Getenv("TOKEN_SECRET")

	b := branca.NewBranca(secret)
	b.SetTTL(uint32(expiry))
	token, err := b.EncodeToString(string(int32(userId)))

	return token, err
}
