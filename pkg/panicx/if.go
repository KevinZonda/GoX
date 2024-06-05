package panicx

func If(exp bool, v any) {
	if exp {
		panic(v)
	}
}
