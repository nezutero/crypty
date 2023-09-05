package hashing

import (
	"crypto/sha256"
	"golang.org/x/crypto/sha3"
)

func CalculateSHA256(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func CalculateSHA3_256(data []byte) []byte {
	hasher := sha3.New256()
	hasher.Write(data)
	return hasher.Sum(nil)
}
