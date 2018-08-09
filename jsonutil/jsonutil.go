package jsonutil

import (
	"bytes"
	"encoding/json"
	"strings"
	"fmt"
	"github.com/lobaro/go-util/reflectutil"
	"strconv"
)

type ByteJsonArray []byte

// Implement json.Marshaler
func (u ByteJsonArray) MarshalJSON() ([]byte, error) {
	var result string
	if u == nil {
		result = "null"
	} else {
		result = strings.Join(strings.Fields(fmt.Sprintf("%d", u)), ",")
	}
	return []byte(result), nil
}

// Implement json.Unmarshaler
func (u *ByteJsonArray) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `[]`)
	tok := strings.Split(str, ",")
	bytes := make([]byte, 0)
	for _, t := range tok {
		i, err := strconv.ParseInt(strings.TrimSpace(t), 0, 10)
		if err != nil {
			return err
		}
		if i < 0 || i > 255 {
			return fmt.Errorf("invalid byte value %d, must be between 0 and 255", i)
		}
		bytes = append(bytes, byte(i))
	}

	*u = ByteJsonArray(bytes)
	return nil
}

type ByteJsonString []byte

func (u ByteJsonString) String() string {
	return string(u)
}

func (u ByteJsonString) MarshalJSON() ([]byte, error) {
	var result string
	if u == nil {
		result = "null"
	} else {
		result = fmt.Sprintf(`"%s"`, u)
	}
	return []byte(result), nil
}

// Implement json.Unmarshaler
func (u *ByteJsonString) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = strings.Trim(str, `"`)
	*u = ByteJsonString(str)
	return nil
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


func MustMarshalIndent(in interface{}) []byte {
	b := &bytes.Buffer{}
	e := json.NewEncoder(b)
	e.SetIndent("", "  ")


	err := e.Encode(in)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}