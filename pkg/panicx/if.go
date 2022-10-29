package panicx

import "fmt"

func PanicIf(exp bool, v any) {
	if exp {
		panic(v)
	}
}

func PanicIfNotNil(b any, v any) {
	if b != nil {
		panic(v)
	}
}

func PanicF(format string, a ...any) {
	panic(fmt.Errorf(format, a...))
}
