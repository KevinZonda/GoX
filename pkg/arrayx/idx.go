package arrayx

func SafeIndex[T any](s []T, i int, ifFailed T) T {
	if i < 0 || i >= len(s) {
		return ifFailed
	}
	return s[i]
}
