package arrayx

func Append[T any](arr []T, items ...T) []T {
	return append(arr, items...)
}

func AppendArr[T any](arr []T, items ...[]T) []T {
	for _, item := range items {
		arr = append(arr, item...)
	}
	return arr
}

func RemoveAt[T any](arr []T, i int) []T {
	if i < 0 || i >= len(arr) {
		return arr
	}
	return append(arr[:i], arr[i+1:]...)
}

func PreAppend[T any](item T, slice []T) []T {
	return append([]T{item}, slice...)
}
