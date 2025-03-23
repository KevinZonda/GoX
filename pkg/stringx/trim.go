package stringx

import "strings"

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func TrimAll(ss []string) []string {
	var r []string
	for _, s := range ss {
		r = append(r, Trim(s))
	}
	return r
}

func TrimAllAndClean(ss []string) []string {
	var r []string
	for _, s := range ss {
		t := Trim(s)
		if t != "" {
			r = append(r, Trim(t))
		}
	}
	return r
}

func TrimLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func TrimUpper(s string) string {
	return strings.ToUpper(strings.TrimSpace(s))
}
