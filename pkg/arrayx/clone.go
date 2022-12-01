package arrayx

func Clone(arr []any) []any {
	newArr := make([]any, len(arr))
	copy(newArr, arr)
	return newArr
}