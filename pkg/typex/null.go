package typex

import (
	"errors"
	"fmt"
)

type Nullable[T any] interface {
	IsNull() bool
	Value() (T, error)
	NotNull()
}

type nullable[T any] struct {
	isNull bool
	value  T
}

func (n nullable[T]) Value() (val T, err error) {
	if n.IsNull() {
		err = errors.New("access null value")
		return
	}
	val = n.value
	return
}

func (n nullable[T]) IsNull() bool {
	return n.isNull
}

func (n nullable[T]) NotNull() {
	if n.IsNull() {
		panic(fmt.Sprintf("Nullable %+v is null however we confirm it should be notnull!", n))
	}
}

func NewNotNull[T any](value T) Nullable[T] {
	return nullable[T]{
		value:  value,
		isNull: false,
	}
}

func NewNull[T any]() Nullable[T] {
	return nullable[T]{
		isNull: true,
	}
}

func NotNull[T any](ns ...Nullable[T]) {
	for _, n := range ns {
		n.NotNull()
	}
}
