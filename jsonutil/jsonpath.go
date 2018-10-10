package jsonutil

import (
	"fmt"
	"strconv"
	"strings"
)

type JsonMap = map[string]interface{}
type JsonArray = []interface{}

func ValueByPath(jsonMap JsonMap, path string) (interface{}, error) {
	if path == "" {
		// Empty path means root element
		return jsonMap, nil
	}

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
			if idx >= int64(len(asArray)) {
				return current, fmt.Errorf("index '%s' >= '%d' out of range", strings.Join(tokens[:i+1], "."), int64(len(asArray)))
			}
			current = asArray[idx]
		} else {
			return current, fmt.Errorf("'%s' must be a JsonMap or JsonArray", strings.Join(tokens[:i+1], "."))
		}

	}

	return current, nil
}
