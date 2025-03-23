package randx

import (
	"crypto/rand"
	"math/big"
)

const ALPHABET_NUMERIC = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const ALPHABET_NUMERIC_NO_ZERO = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
const ALPHABET_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ALPHABET_LOWER = "abcdefghijklmnopqrstuvwxyz"
const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RndId(length int) string {
	return RndIdCharset(length, ALPHABET_NUMERIC)
}

func RndIdCharset(length int, charset string) string {
	if length <= 0 {
		return ""
	}
	charsetLength := big.NewInt(int64(len(charset)))

	b := make([]byte, length)
	for i := range b {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			b[i] = charset[0]
			continue
		}
		b[i] = charset[randomIndex.Int64()]
	}

	return string(b)
}
