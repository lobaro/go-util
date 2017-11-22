package jsonutil

import (
	"encoding/json"
	"reflect"
)

// Takes a unmarshaled json and converts it into another struct
func Convert(in interface{}, out interface{}) error {
	t := reflect.TypeOf(out)
	if t.Kind() != reflect.Ptr {
		panic("out must be pointer but is " + t.Kind().String())
	}

	bytes, err := json.Marshal(in)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &out)
	if err != nil {
		return err
	}
	return nil
}
