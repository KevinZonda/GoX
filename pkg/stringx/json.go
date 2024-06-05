package stringx

import "encoding/json"

func JsonFromObj(o any) string {
	bs, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

func JsonToObj[T any](s string) T {
	var v T
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		panic(err)
	}
	return v
}
