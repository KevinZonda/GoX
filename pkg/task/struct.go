package task

import (
	"github.com/KevinZonda/GoX/pkg/typex"
	"time"
)

type Status int

type Void int

const VoidResult Void = 0

const (
	StatusNone Status = iota
	StatusRunning
	StatusCompleted
)

type Task[T any] struct {
	f      func() T
	status Status
	result *typex.Nullable[T]
}

func New[T any](f func() T) *Task[T] {
	return &Task[T]{
		f:      f,
		status: StatusNone,
		result: typex.NewNull[T](),
	}
}

func (t *Task[T]) Wait() {
	if t.status == StatusCompleted {
		return
	}
	if t.status == StatusRunning {
		for t.status == StatusRunning {
			time.Sleep(50 * time.Millisecond)
		}
	}
	t.status = StatusRunning
	r := t.f()
	t.result.UnNull(r)
	t.status = StatusCompleted
}

func (t *Task[T]) Async() {
	go func() {
		t.Wait()
	}()
}

func (t *Task[T]) IsCompleted() bool {
	return t.status == StatusCompleted
}

func (t *Task[T]) IsRunning() bool {
	return t.status == StatusRunning
}

func (t *Task[T]) Result() (T, error) {
	return t.result.Value()
}

func Wait[T any](t *Task[T]) *Task[T] {
	t.Wait()
	return t
}

func Async[T any](t *Task[T]) *Task[T] {
	t.Async()
	return t
}
