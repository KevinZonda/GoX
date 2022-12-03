package task

import "github.com/KevinZonda/GoX/pkg/typex"

type Status int

type Void int

const (
	StatusNone Status = iota
	StatusRunning
	StatusCompleted
)

type task[T any] struct {
	f      func() T
	status Status
	result *typex.Nullable[T]
}

func NewTask[T any](f func() T) *task[T] {
	return &task[T]{
		f:      f,
		status: StatusNone,
		result: typex.NewNull[T](),
	}
}

func (t *task[T]) Wait() {
	if t.status != StatusNone {
		return
	}
	t.status = StatusRunning
	r := t.f()
	t.result.UnNull(r)
	t.status = StatusCompleted
}

func (t *task[T]) Async() {
	go func() {
		t.Wait()
	}()
}

func (t *task[T]) IsCompleted() bool {
	return t.status == StatusCompleted
}

func (t *task[T]) IsRunning() bool {
	return t.status == StatusRunning
}

func (t *task[T]) Result() (T, error) {
	return t.result.Value()
}

func Wait[T any](t *task[T]) *task[T] {
	t.Wait()
	return t
}

func Async[T any](t *task[T]) *task[T] {
	t.Async()
	return t
}
