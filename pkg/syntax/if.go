package syntax

func IfThenElse[T any](ifCon bool, thenCase T, elseCase T) T {
	if ifCon {
		return thenCase
	}
	return elseCase
}
