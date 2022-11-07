package runex

func IsLowerLetter(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func IsUpperLetter(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func IsLetter(c rune) bool {
	return IsUpperLetter(c) || IsLowerLetter(c)
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

func ForEach[T any](arr []T, f func(index int, data T) T) []T {
	for i, v := range arr {
		arr[i] = f(i, v)
	}
	return arr
}
