package linqx

func From[T any](vs []T) Query[T] {
	return Query[T]{data: vs}
}
