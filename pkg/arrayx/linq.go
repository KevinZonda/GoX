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
	var v []K
	for _, t := range from {
		v = append(v, tx(t))
	}
	return v
}
