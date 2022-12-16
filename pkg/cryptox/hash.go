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

func SHA384(raw []byte) string {
	h := sha512.New384()
	h.Write(raw)
	return hex.EncodeToString(h.Sum(nil))
}

func SHA512_224(raw []byte) string {
	h := sha512.New512_224()
	h.Write(raw)
	return hex.EncodeToString(h.Sum(nil))
}

func SHA512_256(raw []byte) string {
	h := sha512.New512_256()
	h.Write(raw)
	return hex.EncodeToString(h.Sum(nil))
}
