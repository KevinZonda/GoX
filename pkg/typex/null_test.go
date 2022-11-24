package typex

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNullableJsonSerialise(t *testing.T) {
	n := NewNotNull("hello")
	bs, _ := json.Marshal(n)
	if string(bs) != "\"hello\"" {
		t.Fatal("not hello")
	}
	n = NewNull[string]()
	bs, _ = json.Marshal(n)
	if string(bs) != "null" {
		t.Fatal("not null")
	}
}

type mod struct {
	Val Nullable[string]
}

func TestNullableJsonDeserialise(t *testing.T) {
	j := []byte("\"hello\"")
	n := NewNull[string]()
	err := json.Unmarshal(j, &n)
	if err != nil || n.IsNull() {
		t.Fatal(err)
	}
	v, err := n.Value()
	if err != nil || v != "hello" {
		t.Fatal(err)
	}

	j = []byte("{\"Val\": \"hello\"}")
	m := mod{}
	err = json.Unmarshal(j, &m)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", m)
	if !m.Val.IsNull() {
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

func TestNullableNotNull(t *testing.T) {
	n := NewNull[string]()
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("not null")
		}
	}()
	n.NotNull()
}
