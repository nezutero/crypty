package random

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func GenerateRandomHex(length int) (string, error) {
	randomBytes, err := GenerateRandomBytes(length / 2)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(randomBytes), nil
}
