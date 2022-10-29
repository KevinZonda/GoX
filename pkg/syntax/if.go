package syntax

func IfThenElse(ifCon bool, thenCase any, elseCase any) any {
	if ifCon {
		return thenCase
	}
	return elseCase
}
