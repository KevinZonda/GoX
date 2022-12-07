package syntax

// IfThenElse returns the result of the thenCase if the ifCon is true, else the result of the elseCase.
// it provides a ternary operator-like syntax.
func IfThenElse[T any](ifCon bool, thenCase T, elseCase T) T {
	if ifCon {
		return thenCase
	}
	return elseCase
}

// If is a functionalise of if statement
func If(b bool, f func()) {
	if b {
		f()
	}
}
