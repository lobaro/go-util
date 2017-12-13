package reflectutil

import "reflect"

func MustBePointer(i interface{}) {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Ptr {
		panic("value must be pointer but is " + t.Kind().String())
	}
}
