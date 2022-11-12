package typex

type Nullable[T any] struct {
	IsNull bool
	Value  T
}
