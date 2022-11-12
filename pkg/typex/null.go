package typex

type Nullable[T any] struct {
	IsNull bool
	Value  T
}

func NewNotNull[T any](value T) Nullable[T] {
	return Nullable[T]{
		Value:  value,
		IsNull: false,
	}
}

func NewNull[T any]() Nullable[T] {
	return Nullable[T]{
		IsNull: true,
	}
}
