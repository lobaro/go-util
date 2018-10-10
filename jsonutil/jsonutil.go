package jsonutil

import (
	"encoding/json"
	"github.com/lobaro/go-util/reflectutil"
)

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
