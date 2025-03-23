package arrayx

// Last returns the last element of an array.
func Last[T any](l []T) (last T, ok bool) {
	if len(l) == 0 {
		return last, false
	}
	return l[len(l)-1], true
}

// First returns the first element of an array.
func First[T any](l []T) (fst T, ok bool) {
	if len(l) == 0 {
		return fst, false
	}
	return l[0], true
}

// Map returns a new array with the results of calling a provided function on every element in the calling array.
func Map[T any, K any](from []T, tx func(T) K) []K {
	v := make([]K, len(from))
	for i, t := range from {
		v[i] = tx(t)
	}
	return v
}

// ForEach executes a provided function once for each array element.
func ForEach[T any](arr []T, f func(index int, data T)) {
	for i, v := range arr {
		i := i
		v := v
		f(i, v)
	}
}

func ForEachRef[T any](arr []T, f func(index int, data *T)) {
	for i := 0; i < len(arr); i++ {
		f(i, &arr[i])
	}
}

// Any returns true if the provided function returns true for any element in the array.
func Any[T any](vs []T, condition func(T) bool) bool {
	if len(vs) == 0 {
		return false
	}
	for _, v := range vs {
		if condition(v) {
			return true
		}
	}
	return false
}

// All returns true if the provided function returns true for every element in the array.
func All[T any](vs []T, condition func(T) bool) bool {
	if len(vs) == 0 {
		return true
	}
	for _, v := range vs {
		if !condition(v) {
			return false
		}
	}
	return true
}

// Find returns the first element in the array that satisfies the provided testing function.
func Find[T any](vs []T, condition func(T) bool) (val T, ok bool) {
	if len(vs) == 0 {
		return val, false
	}
	for _, v := range vs {
		if condition(v) {
			return v, true
		}
	}
	return val, false
}

// FindIndex returns the index of the first element in the array that satisfies the provided testing function.
func FindIndex[T any](vs []T, condition func(T) bool) (int, bool) {
	if len(vs) == 0 {
		return -1, false
	}
	for i, v := range vs {
		if condition(v) {
			return i, true
		}
	}
	return -1, false
}

// FindLast returns the last element in the array that satisfies the provided testing function.
func FindLast[T any](vs []T, condition func(T) bool) (val T, ok bool) {
	if len(vs) == 0 {
		return val, false
	}
	for i := len(vs) - 1; i >= 0; i-- {
		if condition(vs[i]) {
			return vs[i], true
		}
	}
	return val, false
}

// FindLastIndex returns the index of the last element in the array that satisfies the provided testing function.
func FindLastIndex[T any](vs []T, condition func(T) bool) (int, bool) {
	if len(vs) == 0 {
		return -1, false
	}
	for i := len(vs) - 1; i >= 0; i-- {
		if condition(vs[i]) {
			return i, true
		}
	}
	return -1, false
}

func EditEach[T any](arr []T, f func(index int, data T) T) {
	for i, v := range arr {
		arr[i] = f(i, v)
	}
}

// Where returns a new array with the results of calling a provided function on every element in the array.
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
