package encodingx

import "encoding/json"

func ToJson(v any) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bs), err
}

func ToJsonBytes(v any) ([]byte, error) {
	return json.Marshal(v)
}

func FromJsonTo(from []byte, to any) error {
	return json.Unmarshal(from, to)
}
