package hashing

import (
	"crypto/sha256"
)

func CalculateSHA256(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}
