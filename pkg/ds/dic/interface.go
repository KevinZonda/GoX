package dic

type IDictionary[K comparable, V any] interface {
	Add(key K, value V) bool
	Set(key K, value V)
	Get(key K) (value V, ok bool)
	Remove(key K) (ok bool)
	Contains(key K) bool
	Keys() []K
	Values() []V
	Len() int
	Iterator() map[K]V
}
