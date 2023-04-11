package mapx

func Clone[K comparable, V any](m map[K]V) map[K]V {
	newM := make(map[K]V, len(m))
	for k, v := range m {
		newM[k] = v
	}
	return newM
}

func GetVal[K comparable, V any](m map[string]V, key string, ifNotExists V) V {
	val, ok := m[key]
	if !ok {
		return ifNotExists
	}
	return val
}

// Union if some value has the same key, the latter one will be used
func Union[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	m := make(map[K]V, len(m1))
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		m[k] = v
	}
	return m
}

// Intersection returns the intersection of two maps, the value of the returned map is the value of the first map
func Intersection[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	m := make(map[K]V)
	for k, v := range m1 {
		if _, ok := m2[k]; ok {
			m[k] = v
		}
	}
	return m
}
