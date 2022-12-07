package typex

import "reflect"

// TypeOf returns the type of a type.
func TypeOf[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}

// Is returns true if the type of a type is the same as the type of another type.
func Is[T any](t1 any) bool {
	return reflect.TypeOf(t1) == TypeOf[T]()
}
