package linqx

func (q Query[T]) ForEach(f func(index int, data T) T) Query[T] {
	for i, v := range q.data {
		q.data[i] = f(i, v)
	}
	return q
}
