package cryptox

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func SHA256(raw []byte) string {
	h := sha256.New()
	h.Write(raw)
	return hex.EncodeToString(h.Sum(nil))
}

func SHA512(raw []byte) string {
	h := sha512.New()
	h.Write(raw)
	return hex.EncodeToString(h.Sum(nil))
}
