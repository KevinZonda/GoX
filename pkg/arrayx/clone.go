package arrayx

// Clone returns a copy of the array.
func Clone(arr []any) []any {
	newArr := make([]any, len(arr))
	copy(newArr, arr)
	return newArr
}
