package typex

type Maybe[T any] interface {
	IsNothing() bool
	IsJust() bool
	Value() T
}

type Nothing[T any] struct {
}

func (m *Nothing[T]) IsNothing() bool {
	return true
}

func (m *Nothing[T]) IsJust() bool {
	return false
}

func (m *Nothing[T]) Value() T {
	panic("Get value from Nothing")
}

type Just[T any] struct {
	value T
}

func (m *Just[T]) IsNothing() bool {
	return false
}

func (m *Just[T]) IsJust() bool {
	return true
}

func (m *Just[T]) Value() T {
	return m.value
}

func NewNothing[T any]() *Nothing[T] {
	return &Nothing[T]{}
}

func NewJust[T any](value T) *Just[T] {
	return &Just[T]{
		value: value,
	}
}
