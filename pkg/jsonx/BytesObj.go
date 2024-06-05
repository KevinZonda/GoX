package jsonx

import "encoding/json"

func BytesToObj[T any](s []byte) T {
	var v T
	err := json.Unmarshal(s, &v)
	if err != nil {
		panic(err)
	}
	return v
}

func BytesToObjE[T any](s []byte) (T, error) {
	var v T
	err := json.Unmarshal(s, &v)
	return v, err
}

func BytesToObjPtr[T any](s []byte) *T {
	var v T
	err := json.Unmarshal(s, &v)
	if err != nil || v == nil {
		return nil
	}
	return &v
}

func BytesToObjPtrE[T any](s []byte) (*T, error) {
	var v T
	err := json.Unmarshal(s, &v)
	if err != nil || v == nil {
		return nil, err
	}
	return &v, nil
}

func BytesToMap(s []byte) map[string]any {
	var v map[string]any
	err := json.Unmarshal(s, &v)
	if err != nil {
		panic(err)
	}
	return v
}

func BytesToMapE(s []byte) (map[string]any, error) {
	var v map[string]any
	err := json.Unmarshal(s, &v)
	return v, err
}

func BytesToArr(s []byte) []any {
	var v []any
	err := json.Unmarshal(s, &v)
	if err != nil {
		panic(err)
	}
	return v
}
