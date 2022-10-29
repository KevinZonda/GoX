package arrayx

func HasDuplicate(arr []any) bool {
	m := make(map[any]bool)
	for _, val := range arr {
		_, ok := m[val]
		if ok {
			return true
		}
		m[val] = true
	}
	return false
}
