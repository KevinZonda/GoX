package arrayx

import "testing"

func TestNotNil(t *testing.T) {
	n := NotNil(nil)
	if n {
		t.Error("Should have returned false")
	}
	n = NotNil(1, 2, "x")
	if !n {
		t.Error("Should have returned true")
	}
}
