package jsonutil

import (
	"encoding/json"
	"reflect"
	"strings"
	"fmt"
)

type ByteJsonArray []byte

func (u ByteJsonArray) MarshalJSON() ([]byte, error) {
	var result string
	if u == nil {
		result = "null"
	} else {
		result = strings.Join(strings.Fields(fmt.Sprintf("%d", u)), ",")
	}
	return []byte(result), nil
}

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
