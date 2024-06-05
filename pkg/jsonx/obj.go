package jsonx

func ToObj[T any](s string) T {
	return BytesToObj[T]([]byte(s))
}

func ToObjE[T any](s string) (T, error) {
	return BytesToObjE[T]([]byte(s))
}

func ToObjPtr[T any](s string) *T {
	return BytesToObjPtr[T]([]byte(s))
}

func ToObjPtrE[T any](s string) (*T, error) {
	return BytesToObjPtrE[T]([]byte(s))
}

func ToMap(s string) map[string]any {
	return BytesToMap([]byte(s))
}

func ToMapE(s string) (map[string]any, error) {
	return BytesToMapE([]byte(s))
}

func ToArr(s string) []any {
	return BytesToArr([]byte(s))
}
