package linqx

func (q Query[T]) Where(condition func(T) bool) Query[T] {
	return Query[T]{
		data: Where(q.data, condition),
	}
}

func Where[T1 any](vs []T1, condition func(T1) bool) []T1 {
	if len(vs) == 0 {
		return vs
	}
	var a []T1
	for _, v := range vs {
		if condition(v) {
			a = append(a, v)
		}
	}
	return a
}
