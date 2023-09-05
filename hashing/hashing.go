package hashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	"golang.org/x/crypto/sha3"
)

func CalculateSHA256(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func CalculateSHA224(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	sha256Hash := hasher.Sum(nil)
	sha224Hash := sha256Hash[:28]
	return sha224Hash
}

func CalculateSHA3_256(data []byte) []byte {
	hasher := sha3.New256()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func CalculateSHA1(data []byte) []byte {
	hasher := sha1.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func CalculateMD5(data []byte) []byte {
	hasher := md5.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func CalculateSHA512(data []byte) []byte {
	hasher := sha512.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func CalculateSHA384(data []byte) []byte {
	hasher := sha512.New384()
	hasher.Write(data)
	return hasher.Sum(nil)
}
