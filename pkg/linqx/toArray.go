package linqx

func (q Query[T]) ToArray() []T {
	return q.data
}
