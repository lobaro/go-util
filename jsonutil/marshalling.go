package jsonutil

import (
	"bytes"
	"encoding/json"
)

func MustMarshal(in interface{}) []byte {
	b, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}
	return b
}

func MarshalString(in interface{}) (string, error) {
	b, err := json.Marshal(in)
	return string(b), err
}

func MarshalStringIndent(in interface{}) (string, error) {
	b := &bytes.Buffer{}
	e := json.NewEncoder(b)
	e.SetIndent("", "  ")

	err := e.Encode(in)
	return string(b.Bytes()), err
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

func MustToMap(in string) map[string]interface{} {
	m := make(map[string]interface{})

	err := json.Unmarshal([]byte(in), &m)
	if err != nil {
		panic(err)
	}
	return m
}
