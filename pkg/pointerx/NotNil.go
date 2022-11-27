package pointerx

func NotNil[T any](a ...*any) bool {
	for _, v := range a {
		if v == nil {
			return false
		}
	}
	return true
}
