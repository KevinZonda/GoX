package queue

import "sync"

type ConcurrentQueue[T any] struct {
	l sync.RWMutex
	q ArrQueue[T]
}

func NewConcurrentQueue[T any]() *ConcurrentQueue[T] {
	return &ConcurrentQueue[T]{}
}

func (q *ConcurrentQueue[T]) Enqueue(data T) {
	q.l.Lock()
	defer q.l.Unlock()
	q.q.Enqueue(data)
}

func (q *ConcurrentQueue[T]) Dequeue() (data T, ok bool) {
	q.l.Lock()
	defer q.l.Unlock()
	return q.q.Dequeue()
}

func (q *ConcurrentQueue[T]) TakeFirst() (data T, ok bool) {
	q.l.RLock()
	defer q.l.RUnlock()
	return q.q.TakeFirst()
}

func (q *ConcurrentQueue[T]) ToSlice() []T {
	q.l.RLock()
	defer q.l.RUnlock()
	return q.q.ToSlice()
}
