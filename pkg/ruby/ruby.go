package ruby

import "github.com/KevinZonda/GoX/pkg/panicx"

func Rdr[L any, R any](l L, r R) L {
	return l
}

func Rdl[L any, R any](l L, r R) R {
	return r
}

func Fork[T any](t T) (T, T) {
	return t, t
}

func Pow[T any](t T, n int) []T {
	r := make([]T, n)
	for i := 0; i < n; i++ {
		r[i] = t
	}
	return r
}

func Map[T any, R any](t []T, f func(T) R) []R {
	r := make([]R, len(t))
	for i, v := range t {
		r[i] = f(v)
	}
	return r
}

func RdrErr[L any](l L, r error) L {
	panicx.NotNilErr(r)
	return l
}

func RdlErr[R any](l error, r R) R {
	panicx.NotNilErr(l)
	return r
}

func Apl[T any](t T, list []T) []T {
	return append([]T{t}, list...)
}

func Apr[T any](list []T, t T) []T {
	return append(list, t)
}

func Apply[T any](t []T, f func(T)) {
	for _, v := range t {
		f(v)
	}
}

func ApplyF[T any](t T, f func(T)) {
	f(t)
}

func Left[T any](l T, r T) T {
	return l
}

func Right[T any](l T, r T) T {
	return r
}
