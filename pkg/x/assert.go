package x

func Assert[T any](in any, ifFailed T) T {
	v, ok := in.(T)
	if !ok {
		return ifFailed
	}
	return v
}

func AssertPanic[T any](in any) T {
	v, ok := in.(T)
	if !ok {
		panic("assertion failed")
	}
	return v
}

func AssertAsInt(in any, ifFailed int) int {
	switch vt := in.(type) {
	case int:
		return vt
	case int8:
		return int(vt)
	case int16:
		return int(vt)
	case int32:
		return int(vt)
	case int64:
		return int(vt)
	case uint:
		return int(vt)
	case uint8:
		return int(vt)
	case uint16:
		return int(vt)
	case uint32:
		return int(vt)
	case uint64:
		return int(vt)
	case float32:
		return int(vt)
	case float64:
		return int(vt)
	}
	return ifFailed
}

func AssertAsFloat(in any, ifFailed float64) float64 {
	switch vt := in.(type) {
	case int:
		return float64(vt)
	case int8:
		return float64(vt)
	case int16:
		return float64(vt)
	case int32:
		return float64(vt)
	case int64:
		return float64(vt)
	case uint:
		return float64(vt)
	case uint8:
		return float64(vt)
	case uint16:
		return float64(vt)
	case uint32:
		return float64(vt)
	case uint64:
		return float64(vt)
	case float32:
		return float64(vt)
	case float64:
		return vt
	}
	return ifFailed
}
