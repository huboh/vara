package event

import "reflect"

// getUintPointer returns v's value as a uintptr.
// It panics if v's Kind is not [Chan], [Func], [Map], [Pointer], [Slice], or [UnsafePointer].
func getUintPointer(v any) uintptr {
	return reflect.ValueOf(v).Pointer()
}
