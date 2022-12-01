package encodingx

import (
	"encoding/base32"
	"encoding/base64"
)

type IStrEncoding interface {
	Encode([]byte) (string, error)
	Decode(string) ([]byte, error)
}

func NewBase64Encoder() IStrEncoding {
	return Base64Encoder{}
}

func NewBase32Encoder() IStrEncoding {
	return Base32Encoder{}
}

type Base64Encoder struct{}

func (Base64Encoder) Encode(bs []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(bs), nil
}

func (Base64Encoder) Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

type Base32Encoder struct{}

func (Base32Encoder) Encode(bs []byte) (string, error) {
	return base32.StdEncoding.EncodeToString(bs), nil
}

func (Base32Encoder) Decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}
