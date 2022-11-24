package typex

import (
	"encoding/json"
	"testing"
)

func TestNullableJsonSerialise(t *testing.T) {
	n := NewNotNull("hello")
	bs, _ := json.Marshal(n)
	if string(bs) != "\"hello\"" {
		t.Fatal("not hello")
	}
}

func TestNullableJsonDeserialise(t *testing.T) {
	n := NewNull[string]()
	bs, _ := json.Marshal(n)
	if string(bs) != "null" {
		t.Fatal("not null")
	}
}

func TestNullableIsNull(t *testing.T) {
	n := NewNull[string]()
	if !n.IsNull() {
		t.Fatal("not null")
	}
	n = NewNotNull("hello")
	if n.IsNull() {
		t.Fatal("not null")
	}
}

func TestNullableValue(t *testing.T) {
	n := NewNull[string]()
	_, err := n.Value()
	if err == nil {
		t.Fatal("not null")
	}
	n = NewNotNull("hello")
	v, err := n.Value()
	if err != nil {
		t.Fatal("not null")
	}
	if v != "hello" {
		t.Fatal("not null")
	}
}
