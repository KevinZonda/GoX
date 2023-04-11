package arrayx

// Clone returns a copy of the array.
func Clone[T any](arr []T) []T {
	newArr := make([]T, len(arr))
	copy(newArr, arr)
	return newArr
}
