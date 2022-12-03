package typex

import "testing"

func TestMaybe(t *testing.T) {
	add := func(a, b int) Maybe[int] {
		c := a + b
		if c >= 0 {
			return NewJust(c)
		}
		return NewNothing[int]()
	}
	v := add(1, 2)
	if v.IsNothing() {
		t.Fail()
	}
	if !v.IsJust() {
		t.Fail()
	}
	if v.Value() != 3 {
		t.Fail()
	}
}
