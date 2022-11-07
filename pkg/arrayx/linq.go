package arrayx

func Last[T any](l []T) T {
	if l == nil {
		return nil
	}
	return l[len(l)-1]
}
