package jsonutil

import (
	"encoding/hex"
	"strings"
	"fmt"
)

type HexString []byte

// Implement json.Marshaler
func (u HexString) MarshalJSON() ([]byte, error) {
	var result string
	if u == nil {
		result = "null"
	} else {

		h := hex.EncodeToString(u)
		if len(h) == 0 {
			return []byte(`""`), nil
		}
		result = strings.Join(strings.Fields(fmt.Sprintf(`"0x%s"`, h)), ",")
	}
	return []byte(result), nil
}

func (u HexString)String() string {
	var result string
	if u == nil {
		result = "null"
	} else {

		h := hex.EncodeToString(u)
		if len(h) == 0 {
			return ""
		}
		result = strings.Join(strings.Fields(fmt.Sprintf(`0x%s`, h)), ",")
	}
	return result
}