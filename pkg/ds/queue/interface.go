package queue

type IQueue[T any] interface {
	Enqueue(data T)
	Dequeue() (data T, ok bool)
	TakeFirst() (data T, ok bool)
	ToSlice() []T
}
