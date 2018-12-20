package jsonutil

import (
	"encoding/json"
	"github.com/lobaro/go-util/reflectutil"
)

func IsMap(bytes []byte) bool {
	return len(bytes) > 1 && bytes[0] == '{' && bytes[len(bytes)-1] == '}'
}

// Takes a unmarshaled json and converts it into another struct
func Convert(in interface{}, out interface{}) error {
	reflectutil.MustBePointer(out)

	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &out)
	if err != nil {
		return err
	}
	return nil
}
