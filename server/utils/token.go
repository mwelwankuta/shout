package utils

import (
	"log"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId uint, expiry int64) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userId)),
		ExpiresAt: expiry,
	})

	token, err := claims.SignedString([]byte("secretStuff"))
	log.Printf("After signing %v", token)

	return token, err
}
