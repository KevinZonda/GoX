package x

func Ite[T any](conExpr bool, ifTrue, ifFalse T) T {
	if conExpr {
		return ifTrue
	}
	return ifFalse
}

func IteFunc[T any](conExpr bool, ifTrue, ifFalse func() T) T {
	if conExpr {
		return ifTrue()
	}
	return ifFalse()
}

func IteFuncTrue[T any](conExpr bool, ifTrue func() T, ifFalse T) T {
	if conExpr {
		return ifTrue()
	}
	return ifFalse
}

func IteFuncElse[T any](conExpr bool, ifTrue T, ifFalse func() T) T {
	if conExpr {
		return ifTrue
	}
	return ifFalse()
}
