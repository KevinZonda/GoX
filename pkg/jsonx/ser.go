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

func ToObj[T any](s string) T {
	var v T
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		panic(err)
	}
	return v
}

func ToObjE[T any](s string) (T, error) {
	var v T
	err := json.Unmarshal([]byte(s), &v)
	return v, err
}

func ToObjPtr[T any](s string) *T {
	var v T
	err := json.Unmarshal([]byte(s), &v)
	if err != nil || v == nil {
		return nil
	}
	return &v
}

func ToObjPtrE[T any](s string) (*T, error) {
	var v T
	err := json.Unmarshal([]byte(s), &v)
	if err != nil || v == nil {
		return nil, err
	}
	return &v, nil
}
