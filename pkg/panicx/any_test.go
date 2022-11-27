package panicx

import "testing"

func TestNotNil(t *testing.T) {
	// recovery
	defer func() {
		if r := recover(); r != nil {
			t.Log("Recovered from", r)
		}
	}()
	NotNil(nil)
	t.Error("Should have panicked")
}

func TestNotNil2(t *testing.T) {
	// recovery
	defer func() {
		if r := recover(); r != nil {
			t.Log("Recovered from", r)
			t.Error("Should have not panicked")
		}
	}()
	NotNil(1, 2, "x")
}
