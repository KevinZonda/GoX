package queue

import "sync"

type ConcurrentQueue[T any] struct {
	sync.RWMutex
	ArrQueue[T]
}

func NewConcurrentQueue[T any]() *ConcurrentQueue[T] {
	return &ConcurrentQueue[T]{}
}

func (q *ConcurrentQueue[T]) Enqueue(data T) {
	q.Lock()
	defer q.Unlock()
	q.ArrQueue.Enqueue(data)
}

func (q *ConcurrentQueue[T]) Dequeue() (data T, ok bool) {
	q.Lock()
	defer q.Unlock()
	return q.ArrQueue.Dequeue()
}

func (q *ConcurrentQueue[T]) TakeFirst() (data T, ok bool) {
	q.RLock()
	defer q.RUnlock()
	return q.ArrQueue.TakeFirst()
}

func (q *ConcurrentQueue[T]) ToSlice() []T {
	q.RLock()
	defer q.RUnlock()
	return q.ArrQueue.ToSlice()
}