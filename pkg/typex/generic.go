package typex

import "reflect"

func TypeOf[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}

func Is[T any](t1 any) bool {
	return reflect.TypeOf(t1) == TypeOf[T]()
}
