package panicx

import "strconv"

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
