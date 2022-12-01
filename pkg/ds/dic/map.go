package dic

import "fmt"

type Dictionary[K comparable, V any] struct {
	fields map[K]V
}

func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{
		fields: make(map[K]V),
	}
}

func (d *Dictionary[K, V]) Add(key K, value V) bool {
	_, ok := d.fields[key]
	if ok {
		return false
	}
	d.fields[key] = value
	return true
}

func (d *Dictionary[K, V]) Set(key K, value V) {
	d.fields[key] = value
}

func (d *Dictionary[K, V]) Get(key K) (value V, ok bool) {
	value, ok = d.fields[key]
	return
}

func (d *Dictionary[K, V]) Remove(key K) bool {
	_, ok := d.fields[key]
	if ok {
		delete(d.fields, key)
		return true
	}
	return false
}

func (d *Dictionary[K, V]) Contains(key K) bool {
	_, ok := d.fields[key]
	return ok
}

func (d *Dictionary[K, V]) Len() int {
	return len(d.fields)
}

func (d *Dictionary[K, V]) Keys() []K {
	keys := make([]K, 0, len(d.fields))
	for k := range d.fields {
		keys = append(keys, k)
	}
	return keys
}

func (d *Dictionary[K, V]) Values() []V {
	values := make([]V, 0, len(d.fields))
	for _, v := range d.fields {
		values = append(values, v)
	}
	return values
}

func (d *Dictionary[K, V]) String() string {
	return fmt.Sprintf("%v", d.fields)
}

func (d *Dictionary[K, V]) CloneToMap() map[K]V {
	m := make(map[K]V, len(d.fields))
	for k, v := range d.fields {
		m[k] = v
	}
	return m
}

func (d *Dictionary[K, V]) Clone() *Dictionary[K, V] {
	return &Dictionary[K, V]{
		fields: d.CloneToMap(),
	}
}
