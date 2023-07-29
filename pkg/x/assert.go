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
	switch in.(type) {
	case int:
		return in.(int)
	case int8:
		return int(in.(int8))
	case int16:
		return int(in.(int16))
	case int32:
		return int(in.(int32))
	case int64:
		return int(in.(int64))
	case uint:
		return int(in.(uint))
	case uint8:
		return int(in.(uint8))
	case uint16:
		return int(in.(uint16))
	case uint32:
		return int(in.(uint32))
	case uint64:
		return int(in.(uint64))
	case float32:
		return int(in.(float32))
	case float64:
		return int(in.(float64))
	}
	return ifFailed
}

func AssertAsFloat(in any, ifFailed float64) float64 {
	switch in.(type) {
	case int:
		return float64(in.(int))
	case int8:
		return float64(in.(int8))
	case int16:
		return float64(in.(int16))
	case int32:
		return float64(in.(int32))
	case int64:
		return float64(in.(int64))
	case uint:
		return float64(in.(uint))
	case uint8:
		return float64(in.(uint8))
	case uint16:
		return float64(in.(uint16))
	case uint32:
		return float64(in.(uint32))
	case uint64:
		return float64(in.(uint64))
	case float32:
		return float64(in.(float32))
	case float64:
		return in.(float64)
	}
	return ifFailed
}
