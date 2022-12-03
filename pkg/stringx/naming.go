package stringx

import (
	"regexp"
	"strings"
)

func SnakeCase(input string) string {
	var matchChars = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAlpha = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchChars.ReplaceAllString(input, "${1}_${2}")
	snake = matchAlpha.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
