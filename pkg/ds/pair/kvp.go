package pair

import "fmt"

type KeyValuePair[K, V any] struct {
	key   K
	value V
}

func NewKeyValuePair[K, V any](key K, value V) KeyValuePair[K, V] {
	return KeyValuePair[K, V]{key: key, value: value}
}

func (kvp KeyValuePair[K, V]) Key() K {
	return kvp.key
}

func (kvp KeyValuePair[K, V]) Value() V {
	return kvp.value
}

func (kvp KeyValuePair[K, V]) SetKey(key K) {
	kvp.key = key
}

func (kvp KeyValuePair[K, V]) SetValue(value V) {
	kvp.value = value
}

func (kvp KeyValuePair[K, V]) String() string {
	return fmt.Sprintf("%v=%v", kvp.key, kvp.value)
}
