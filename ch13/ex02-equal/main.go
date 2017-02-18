package main

import (
	"reflect"
	"unsafe"
)

func HasCycle(x interface{}) bool {
	seen := make(map[unsafe.Pointer]bool)
	return hasCycle(reflect.ValueOf(x), seen)
}

func hasCycle(x reflect.Value, seen map[unsafe.Pointer]bool) bool {
	if !x.IsValid() {
		return false
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		if isSeen(x, seen) {
			return true
		}
		return hasCycle(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if hasCycle(x.Index(i), seen) {
				return true
			}
		}

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if hasCycle(x.Field(i), seen) {
				return true
			}
		}

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if hasCycle(k, seen) {
				return true
			}
			if hasCycle(x.MapIndex(k), seen) {
				return true
			}
		}
	}
	return false
}

func isSeen(x reflect.Value, seen map[unsafe.Pointer]bool) bool {
	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		if seen[xptr] {
			return true // already seen
		}
		seen[xptr] = true
	}
	return false
}
