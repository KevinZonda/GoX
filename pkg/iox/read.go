package iox

import (
	"io"
	"os"
	"strings"
)

func ReadAllByte(file string) (bs []byte, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

func ReadAllText(file string) (str string, err error) {
	var bs []byte
	bs, err = ReadAllByte(file)
	if err != nil {
		return
	}
	str = string(bs)
	return
}

func ReadAllLines(file string) (strs []string, err error) {
	var str string
	str, err = ReadAllText(file)
	if err != nil {
		return
	}
	strs = strings.Split(str, "\n")
	return
}

func IOReadAllText(i io.Reader) (string, error) {
	bs, err := io.ReadAll(i)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func IOReadAllLines(i io.Reader) ([]string, error) {
	str, err := IOReadAllText(i)
	if err != nil {
		return nil, err
	}
	return strings.Split(str, "\n"), nil
}
