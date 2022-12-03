package typex

import (
	"testing"
)

func TestMultiType(t *testing.T) {
	m := NewMultiType2WithT1[double, string](1)
	if !m.IsT1() {
		t.Errorf("IsT1 failed")
		t.Fail()
	}
	if m.IsT2() {
		t.Errorf("IsT2 failed")
		t.Fail()
	}
	m.SetT2("hello")
	if !m.IsT2() {
		t.Errorf("IsT2 failed")
		t.Fail()
	}
	if m.IsT1() {
		t.Errorf("IsT1 failed")
		t.Fail()
	}
	if m.GetT2() != "hello" {
		t.Errorf("GetT2 failed")
		t.Fail()
	}
}
