package arrayx

import (
	"github.com/KevinZonda/GoX/pkg/constrains"
	"math/rand"
	"time"
)

func HasDuplicate[T comparable](arr []T) bool {
	m := make(map[T]bool)
	for _, val := range arr {
		_, ok := m[val]
		if ok {
			return true
		}
		m[val] = true
	}
	return false
}

func RemoveDuplicate[T comparable](arr []T) []T {
	m := make(map[T]bool)
	var a []T
	for _, val := range arr {
		_, ok := m[val]
		if ok {
			continue
		}
		m[val] = true
		a = append(a, val)
	}
	return a
}

func Shuffle[T any](array []T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(array) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

func Reverse[T any](a []T) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		pos := len(a) - 1 - i
		a[i], a[pos] = a[pos], a[i]
	}
}

func Sum[T constrains.Calculable](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}
