package errorx

func IgnoreErr[T any](val T, _ error) T {
	return val
}
