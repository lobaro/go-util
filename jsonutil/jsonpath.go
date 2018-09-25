package jsonutil

import (
	"fmt"
	"strconv"
	"strings"
)

type JsonMap = map[string]interface{}
type JsonArray = []interface{}

func ValueByPath(jsonMap JsonMap, path string) (interface{}, error) {
	tokens := strings.Split(path, ".")
	var current interface{}
	current = jsonMap
	for i, t := range tokens {
		asMap, isMap := current.(JsonMap)
		asArray, isArray := current.(JsonArray)
		if isMap {
			current = asMap[t]
		} else if isArray {
			idx, err := strconv.ParseInt(t, 10, 64)
			if err != nil {
				return current, fmt.Errorf("token '%s' must be an array index", t)
			}
			if int64(len(asArray)) >= idx {
				return current, fmt.Errorf("index '%s' out of range", strings.Join(tokens[:i+1], "."))
			}
			current = asArray[idx]
		} else {
			return current, fmt.Errorf("'%s' must be a JsonMap or JsonArray", strings.Join(tokens[:i+1], "."))
		}

	}

	return current, nil
}
