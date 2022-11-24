package arrayx

import "testing"

func TestFirst(t *testing.T) {
	v := []int{1, 2, 3}
	fst, ok := First(v)
	if !ok || fst != 1 {
		t.Fail()
	}
	var v2 []int
	fst, ok = First(v2)
	if ok {
		t.Fail()
	}
}

func TestLast(t *testing.T) {
	v := []int{1, 2, 3}
	last, ok := Last(v)
	if !ok || last != 3 {
		t.Fail()
	}
	var v2 []int
	last, ok = Last(v2)
	if ok {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	v := []int{1, 2, 3}
	mapped := Map(v, func(i int) int {
		return i * 2
	})
	if len(mapped) != 3 || mapped[0] != 2 || mapped[1] != 4 || mapped[2] != 6 {
		t.Fail()
	}
}
