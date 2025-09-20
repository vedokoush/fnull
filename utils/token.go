package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)[:length]
}