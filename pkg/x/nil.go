package x

func NotNil(a ...any) bool {
	for _, v := range a {
		if v == nil {
			return false
		}
	}
	return true
}

func PanicIfNil(a ...any) {
	for _, v := range a {
		if v == nil {
			panic("nil")
		}
	}
}
