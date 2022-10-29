package encodingx

import "encoding/base64"

func ToBase64(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}

func StringToBase64(s string) string {
	return ToBase64([]byte(s))
}

func FromBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func FromBase64ToString(s string) (string, error) {
	bs, err := FromBase64(s)
	if err != nil {
		return "", err
	}
	return string(bs), err
}
