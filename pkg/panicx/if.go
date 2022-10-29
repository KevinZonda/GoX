package panicx

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
