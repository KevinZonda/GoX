package typex

import "reflect"

type MultiType2[T1, T2 any] struct {
	obj any
	t   reflect.Type
}

func NewMultiType2WithT1[T1, T2 any](obj T1) *MultiType2[T1, T2] {
	return &MultiType2[T1, T2]{
		obj: obj,
		t:   reflect.TypeOf(obj),
	}
}

func NewMultiType2WithT2[T1, T2 any](obj T2) *MultiType2[T1, T2] {
	return &MultiType2[T1, T2]{
		obj: obj,
		t:   reflect.TypeOf(obj),
	}
}

func (m *MultiType2[T1, T2]) IsT1() bool {
	return m.t == reflect.TypeOf(m.obj)
}

func (m *MultiType2[T1, T2]) IsT2() bool {
	return m.t == reflect.TypeOf(m.obj)
}

func (m *MultiType2[T1, T2]) GetT1() T1 {
	return m.obj.(T1)
}

func (m *MultiType2[T1, T2]) GetT2() T2 {
	return m.obj.(T2)
}

func (m *MultiType2[T1, T2]) SetT1(t T1) {
	m.obj = t
}

func (m *MultiType2[T1, T2]) SetT2(t T2) {
	m.obj = t
}

func (m *MultiType2[T1, T2]) Get() any {
	return m.obj
}

type MultiType3[T1, T2, T3 any] struct {
	obj any
	t   reflect.Type
}

func NewMultiType3WithT1[T1, T2, T3 any](obj T1) *MultiType3[T1, T2, T3] {
	return &MultiType3[T1, T2, T3]{
		obj: obj,
		t:   reflect.TypeOf(obj),
	}
}

func NewMultiType3WithT2[T1, T2, T3 any](obj T2) *MultiType3[T1, T2, T3] {
	return &MultiType3[T1, T2, T3]{
		obj: obj,
		t:   reflect.TypeOf(obj),
	}
}

func NewMultiType3WithT3[T1, T2, T3 any](obj T3) *MultiType3[T1, T2, T3] {
	return &MultiType3[T1, T2, T3]{
		obj: obj,
		t:   reflect.TypeOf(obj),
	}
}

func (m *MultiType3[T1, T2, T3]) IsT1() bool {
	return m.t == reflect.TypeOf(m.obj)
}

func (m *MultiType3[T1, T2, T3]) IsT2() bool {
	return m.t == reflect.TypeOf(m.obj)
}

func (m *MultiType3[T1, T2, T3]) IsT3() bool {
	return m.t == reflect.TypeOf(m.obj)
}

func (m *MultiType3[T1, T2, T3]) GetT1() T1 {
	return m.obj.(T1)
}

func (m *MultiType3[T1, T2, T3]) GetT2() T2 {
	return m.obj.(T2)
}

func (m *MultiType3[T1, T2, T3]) GetT3() T3 {
	return m.obj.(T3)
}

func (m *MultiType3[T1, T2, T3]) SetT1(t T1) {
	m.obj = t
}

func (m *MultiType3[T1, T2, T3]) SetT2(t T2) {
	m.obj = t
}

func (m *MultiType3[T1, T2, T3]) SetT3(t T3) {
	m.obj = t
}

func (m *MultiType3[T1, T2, T3]) Get() any {
	return m.obj
}
