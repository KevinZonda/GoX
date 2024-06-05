package stringx

import "strings"

func Join(sep string, ss ...string) string {
	return strings.Join(ss, sep)
}

func JoinNotEmpty(sep string, ss ...string) string {
	var _ss []string
	for _, s := range ss {
		if s != "" {
			_ss = append(_ss, s)
		}
	}
	return strings.Join(_ss, sep)
}
