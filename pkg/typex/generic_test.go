package typex

import (
	"reflect"
	"testing"
)

func TestIs(t *testing.T) {
	if !Is[int](1) {
		t.Errorf("Is[int] failed")
		t.Fail()
	}
	if Is[int]("hello") {
		t.Errorf("Is[int] failed")
		t.Fail()
	}
}

func TestTypeOf(t *testing.T) {
	if TypeOf[int]().Kind() != reflect.Int {
		t.Errorf("TypeOf[int] failed")
		t.Fail()
	}
}
