package runex

func IsLowerLetter(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func IsUpperLetter(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func IsLetter(c rune) bool {
	return IsUpperLetter(c) || IsLowerLetter(c)
}
