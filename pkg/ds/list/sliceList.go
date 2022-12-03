package list

func increment(cur int) int {
	if cur < 1001 {
		return cur * 2
	}
	return cur + 1000
}

type SliceList[T comparable] struct {
	slice []T
	count int
	cap   int
}

func NewSliceList[T comparable]() *SliceList[T] {
	return &SliceList[T]{
		slice: make([]T, 10),
		count: 0,
		cap:   10,
	}
}

func (l *SliceList[T]) Add(i T) {
	if l.count == l.cap {
		newCap := increment(l.cap)
		newSlice := make([]T, newCap)
		copy(newSlice, l.slice)
		l.slice = newSlice
	}
	l.slice[l.count] = i
	l.count++
}

func (l *SliceList[T]) AddAll(i []T) {
	for _, v := range i {
		l.Add(v)
	}
}

func (l *SliceList[T]) Count() int {
	return l.count
}

func (l *SliceList[T]) Contains(i T) bool {
	for _, v := range l.slice {
		if v == i {
			return true
		}
	}
	return false
}

func (l *SliceList[T]) ContainsAll(i []T) bool {
	for _, v := range i {
		if !l.Contains(v) {
			return false
		}
	}
	return true
}

func (l *SliceList[T]) IndexOf(e T) int {
	for i := 0; i < l.count; i++ {
		if l.slice[i] == e {
			return i
		}
	}
	return -1
}

func (l *SliceList[T]) Remove(i T) bool {
	// TODO:
	return false
}

func (l *SliceList[T]) RemoveAt(i int) bool {
	// TODO:
	return false
}

func (l *SliceList[T]) Clear() {
	l.count = 0
	l.slice = make([]T, 10)
	l.cap = 10
	return
}

func (l *SliceList[T]) ToSlice() []T {
	return l.slice[:l.count]
}
