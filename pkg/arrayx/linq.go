package arrayx

func Last[T any](l []T) T {
	if l == nil {
		return nil
	}
	return l[len(l)-1]
}

func Map[T any, K any](from []T, tx func(T) K) []K {
	var v []K
	for _, t := range from {
		v = append(v, tx(t))
	}
	return v
}
