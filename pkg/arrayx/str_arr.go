package arrayx

func WhereNotEmpty(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	return Where(arr, func(s string) bool {
		return s != ""
	})
}

func WhereNotNil[T any](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}
	return Where[T](arr, func(v T) bool {
		return v != nil
	})
}
