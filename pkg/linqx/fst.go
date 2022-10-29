package linqx

func (q Query[T]) First() T {
	return q.data[0]
}

func (q Query[T]) Last() T {
	return q.data[len(q.data)-1]
}

func (q Query[T]) Second() T {
	return q.data[1]
}

func (q Query[T]) Count() int {
	return len(q.data)
}
