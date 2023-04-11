package arrayx

// Contains returns true if the array contains the item.
func Contains[T comparable](arr []T, item T) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func ContainsAny[T comparable](arr []T, items ...T) bool {
	for _, item := range items {
		if Contains(arr, item) {
			return true
		}
	}
	return false
}
