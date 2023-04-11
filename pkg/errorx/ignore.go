package errorx

import (
	"errors"
	"strings"
)

func IgnoreErr[T any](val T, _ error) T {
	return val
}

func Wrap(err error, tag string) error {
	if err == nil {
		return nil
	}
	return errors.New("<" + tag + ">" + err.Error() + "</" + tag + ">")
}

func IsWrapped(err error, tag string) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	if !strings.HasSuffix(msg, "</"+tag+">") {
		return false
	}
	if !strings.HasPrefix(msg, "<"+tag+">") {
		return false
	}
	return true
}

func Unwrap(err error, tag string) string {
	if err == nil {
		return ""
	}
	if !IsWrapped(err, tag) {
		return ""
	}
	msg := err.Error()
	msg = msg[len("<"+tag+">"):]
	msg = msg[:len(msg)-len("</"+tag+">")]
	return msg
}
