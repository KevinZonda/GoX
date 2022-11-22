package typex

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Enqueue(data T) {
	q.data = append(q.data, data)
}

func (q *Queue[T]) Dequeue() (data T, ok bool) {
	data, ok = q.TakeFirst()
	if ok {
		q.data = q.data[1:]
	}
	return
}

func (q *Queue[T]) TakeFirst() (data T, ok bool) {
	if len(q.data) == 0 {
		return data, false
	}
	data = q.data[0]
	return data, true
}

func (q *Queue[T]) ToSlice() []T {
	return q.data
}
