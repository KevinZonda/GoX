package stringx

import "strings"

func Concat(sep string, ss ...string) string {
	var sb strings.Builder
	for i, s := range ss {
		s := s
		if i > 0 {
			sb.WriteString(sep)
		}
		sb.WriteString(s)
	}
	return sb.String()
}

func ConcatIfNotEmpty(sep string, ss ...string) string {
	var sb strings.Builder
	for i, s := range ss {
		s := strings.TrimSpace(s)
		if s != "" {
			if i > 0 {
				sb.WriteString(sep)
			}
			sb.WriteString(s)
		}
	}
	return sb.String()
}
