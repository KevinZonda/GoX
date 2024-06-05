package jsonx

import "encoding/json"

func ToString(o any) string {
	bs, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

func ToStringE(o any) (string, error) {
	bs, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// ToStringI with indent
func ToStringI(o any) string {
	bs, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(bs)
}

// ToStringIE with indent
func ToStringIE(o any) (string, error) {
	bs, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
