package utils

import (
	"crypto/rand"
	"log"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

func InitRandomToken(length int) string {
	rand, err := random(length)
	if err != nil {
		log.Fatalln(err)
	}
	return rand
}

func random(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}
