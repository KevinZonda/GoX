package typex

import "strings"

type char rune

func ToCharArray(s string) []char {
	return []char(s)
}

func Length(s string) int {
	return len(ToCharArray(s))
}

func ToString(cs []char) string {
	return string(cs)
}

func (c char) ToString() string {
	return string(c)
}

func (c char) ToRune() rune {
	return rune(c)
}

func (c char) Append(s string) string {
	var sb strings.Builder
	sb.WriteRune(rune(c))
	sb.WriteString(s)
	return sb.String()
}

func (c char) Prepend(s string) string {
	var sb strings.Builder
	sb.WriteString(s)
	sb.WriteRune(rune(c))
	return sb.String()
}

func ToRunes(cs []char) []rune {
	var m any = cs
	return m.([]rune)
}
