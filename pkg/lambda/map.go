package lambda

func Map[T any, R any](in []T, f func(T) R) []R {
	out := make([]R, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func ForEach[T any](in []T, f func(T)) {
	for _, v := range in {
		f(v)
	}
}

func Filter[T any](in []T, f func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func Find[T any](in []T, f func(T) bool) (T, bool) {
	for _, v := range in {
		if f(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

func Any[T any](in []T, f func(T) bool) bool {
	for _, v := range in {
		if f(v) {
			return true
		}
	}
	return false
}

func All[T any](in []T, f func(T) bool) bool {
	for _, v := range in {
		if !f(v) {
			return false
		}
	}
	return true
}
