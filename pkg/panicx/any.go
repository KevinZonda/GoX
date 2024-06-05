package panicx

import (
	"fmt"
	"strconv"
)

func Panic(a any) {
	panic(a)
}

func NotNil(a ...any) {
	for i, v := range a {
		if v == nil {
			panic("Nil pointer at index " + strconv.Itoa(i))
		}
	}
}

func NotNilErr(e error) {
	if e != nil {
		panic(e)
	}
}

func NotImplemented() {
	panic("Not implemented")
}

func PanicF(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}

func Format(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}
