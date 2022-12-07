package dic

import (
	"sync"
)

// ConcurrentDictionary is a thread-safe dictionary.
type ConcurrentDictionary[K comparable, V any] struct {
	m *Dictionary[K, V]
	l sync.RWMutex
}

// NewConcurrentDictionary creates a new ConcurrentDictionary.
func NewConcurrentDictionary[K comparable, V any]() *ConcurrentDictionary[K, V] {
	return &ConcurrentDictionary[K, V]{
		m: NewDictionary[K, V](),
	}
}

// Add adds a new item to the dictionary.
func (d *ConcurrentDictionary[K, V]) Add(key K, value V) bool {
	d.l.Lock()
	defer d.l.Unlock()
	return d.m.Add(key, value)
}

func (d *ConcurrentDictionary[K, V]) Set(key K, value V) {
	d.l.Lock()
	defer d.l.Unlock()
	d.m.Set(key, value)
}

func (d *ConcurrentDictionary[K, V]) Get(key K) (V, bool) {
	d.l.RLock()
	defer d.l.RUnlock()
	return d.m.Get(key)
}

func (d *ConcurrentDictionary[K, V]) Remove(key K) bool {
	d.l.Lock()
	defer d.l.Unlock()
	return d.m.Remove(key)
}

func (d *ConcurrentDictionary[K, V]) Contains(key K) bool {
	d.l.RLock()
	defer d.l.RUnlock()
	return d.m.Contains(key)
}

func (d *ConcurrentDictionary[K, V]) Count() int {
	d.l.RLock()
	defer d.l.RUnlock()
	return d.m.Count()
}

func (d *ConcurrentDictionary[K, V]) Keys() []K {
	d.l.RLock()
	defer d.l.RUnlock()
	return d.m.Keys()
}

func (d *ConcurrentDictionary[K, V]) Values() []V {
	d.l.RLock()
	defer d.l.RUnlock()
	return d.m.Values()
}

func (d *ConcurrentDictionary[K, V]) Iterator() map[K]V {
	return d.m.Iterator()
}

func (d *ConcurrentDictionary[K, V]) MarshalJSON() ([]byte, error) {
	d.l.RLock()
	defer d.l.RUnlock()
	return d.m.MarshalJSON()
}

func (d *ConcurrentDictionary[K, V]) UnmarshalJSON(b []byte) error {
	d.l.Lock()
	defer d.l.Unlock()
	return d.m.UnmarshalJSON(b)
}
