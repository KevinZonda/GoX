package typex

type Maybe[T any] interface {
	IsNothing() bool
	IsJust() bool
	Value() (value T, ok bool)
}

type Nothing[T any] struct {
}

func (m *Nothing[T]) IsNothing() bool {
	return true
}

func (m *Nothing[T]) IsJust() bool {
	return false
}

func (m *Nothing[T]) Value() (value T, ok bool) {
	return value, false
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

func (m *Just[T]) Value() (value T, ok bool) {
	return m.value, true
}
