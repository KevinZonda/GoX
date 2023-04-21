package arrayx

func NotNil(a ...any) bool {
	for _, v := range a {
		if v == nil {
			return false
		}
	}
	return true
}

func NotEmpty[T any](a ...[]T) bool {
	for _, v := range a {
		if len(v) == 0 {
			return false
		}
	}
	return true
}
