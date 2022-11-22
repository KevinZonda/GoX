package httpx

import (
	"io"
	"io/ioutil"
	"net/http"
)

func GET(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
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
