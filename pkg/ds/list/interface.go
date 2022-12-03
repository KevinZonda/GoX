package list

type IList[T any] interface {
	Add(i T)
	AddAll(i []T)
	Contains(i T) bool
	IndexOf(i T) int
	Remove(i T) bool
	RemoveAt(i int) bool
	ToSlice() []T
	Count() int
}
