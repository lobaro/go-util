package reflectutil

import (
	"reflect"
	"runtime/debug"
)

func MustBePointer(i interface{}) {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Ptr {
		debug.PrintStack()
		panic("panic: value must be pointer but is " + t.Kind().String())
	}
}
