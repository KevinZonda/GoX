package stringx

import "strings"

func ReplaceAll(origin string, newVal string, oldVals ...string) string {
	for _, oldVal := range oldVals {
		origin = strings.ReplaceAll(origin, oldVal, newVal)
	}
	return origin
}
