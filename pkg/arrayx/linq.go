package arrayx

func Last[T any](l []T) (last T, ok bool) {
	if l == nil || len(l) == 0 {
		return last, false
	}
	return l[len(l)-1], true
}

func First[T any](l []T) (fst T, ok bool) {
	if l == nil || len(l) == 0 {
		return fst, false
	}
	return l[0], true
}

func Map[T any, K any](from []T, tx func(T) K) []K {
	v := make([]K, len(from))
	for i, t := range from {
		v[i] = tx(t)
	}
	return v
}

func ForEach[T any](arr []T, f func(index int, data T)) {
	for i, v := range arr {
		i := i
		v := v
		f(i, v)
	}
}

func EditAll[T any](arr []T, f func(index int, data T) T) {
	for i, v := range arr {
		arr[i] = f(i, v)
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
