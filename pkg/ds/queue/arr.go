package queue

type ArrQueue[T any] struct {
	data []T
}

func (q *ArrQueue[T]) Enqueue(data T) {
	q.data = append(q.data, data)
}

func NewArrQueue[T any]() *ArrQueue[T] {
	return &ArrQueue[T]{}
}

func (q *ArrQueue[T]) Dequeue() (data T, ok bool) {
	data, ok = q.TakeFirst()
	if ok {
		q.data = q.data[1:]
	}
	return
}

func (q *ArrQueue[T]) TakeFirst() (data T, ok bool) {
	if len(q.data) == 0 {
		return data, false
	}
	data = q.data[0]
	return data, true
}

func (q *ArrQueue[T]) ToSlice() []T {
	return q.data
}
