package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte("super-secret-key")

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  username,
		"ExpiresAt": time.Now().Add(24 * time.Hour),
	})
	return token.SignedString(SECRET_KEY)
}
