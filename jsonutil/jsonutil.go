package jsonutil

import (
	"encoding/json"
	"strings"
	"fmt"
	"github.com/Lobaro/go-util/reflectutil"
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

type ByteJsonString []byte

func (u ByteJsonString) MarshalJSON() ([]byte, error) {
	var result string
	if u == nil {
		result = "null"
	} else {
		result = fmt.Sprintf(`"%s"`, u)
	}
	return []byte(result), nil
}

// Takes a unmarshaled json and converts it into another struct
func Convert(in interface{}, out interface{}) error {
	reflectutil.MustBePointer(out)

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

func MustMarshal(in interface{}) []byte {
	bytes, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}
	return bytes
}
