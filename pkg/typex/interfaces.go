package typex

type ICloneable[T any] interface {
	Clone() T
}
