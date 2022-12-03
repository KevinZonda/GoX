package typex

import (
	"encoding/json"
	"errors"
	"fmt"
)

type INullable[T any] interface {
	IsNull() bool
	Value() (T, error)
	NotNull()
	Null()
	UnNull(value T)
}

type Nullable[T any] struct {
	isNull bool
	value  T
}

func (n *Nullable[T]) Value() (val T, err error) {
	if n.IsNull() {
		err = errors.New("access null value")
		return
	}
	val = n.value
	return
}

func (n *Nullable[T]) IsNull() bool {
	if n == nil {
		return true
	}
	return n.isNull
}

func (n *Nullable[T]) NotNull() {
	if n.IsNull() {
		panic(fmt.Sprintf("Nullable %+v is null however we confirm it should be notnull!", n))
	}
}

func (n *Nullable[T]) Null() {
	n.isNull = true
}

func (n *Nullable[T]) UnNull(value T) {
	n.isNull = false
	n.value = value
}

func (n *Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.isNull {
		return json.Marshal(nil)
	}
	return json.Marshal(n.value)
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	var m *T = nil
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	n.isNull = m == nil
	if m != nil {
		n.value = *m
	}
	return nil
}

func NewNotNull[T any](value T) *Nullable[T] {
	return &Nullable[T]{
		value:  value,
		isNull: false,
	}
}

func NewNull[T any]() *Nullable[T] {
	return &Nullable[T]{
		isNull: true,
	}
}

func NotNull[T any](ns ...Nullable[T]) {
	for _, n := range ns {
		n.NotNull()
	}
}
