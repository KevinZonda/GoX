package asyncx

type Task[T any] struct {
	Task      func() T
	isRunning bool
	endCh     chan T
	rst       *T
}

func NewTask[T any](task func() T) *Task[T] {
	return &Task[T]{
		Task:      task,
		isRunning: false,
		endCh:     nil,
		rst:       nil,
	}
}

func (a *Task[T]) IsRunning() bool {
	return a.isRunning
}

func (a *Task[T]) IsCompleted() bool {
	return !a.isRunning && a.rst != nil
}

func (a *Task[T]) Result() *T {
	return a.rst
}

func (a *Task[T]) Async() {
	if a.isRunning {
		return
	}
	a.isRunning = true
	a.endCh = make(chan T)
	go func() {
		a.Task()
		rst := a.Task()
		a.rst = &rst
		a.endCh <- rst
		a.isRunning = false
	}()
}

func (a *Task[T]) wait() {
	if !a.isRunning {
		return
	}
	a.isRunning = true
	rst := a.Task()
	a.rst = &rst
	a.isRunning = false
}

func (a *Task[T]) Wait() T {
	if !a.isRunning {
		return a.Task()
	}
	return <-a.endCh
}
