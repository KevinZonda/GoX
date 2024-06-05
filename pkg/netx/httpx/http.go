package httpx

import (
	"encoding/json"
	"io"
	"net/http"
)

func GETString(url string) (string, error) {
	bs, err := GETBytes(url)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func GETBytes(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err = io.ReadAll(resp.Body)
	return
}

func GETJson[T any](url string) (T, error) {
	var v T
	bs, err := GETBytes(url)
	if err != nil {
		return v, err
	}
	err = json.Unmarshal(bs, &v)
	return v, err
}

var _defaultUA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"

func SetDefaultUA(ua string) {
	_defaultUA = ua
}

func GetDefaultUA() string {
	return _defaultUA
}

func NewGetJsonRequest(url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", _defaultUA)
	req.Header.Set("Accept", "application/json")
	return req, nil
}
