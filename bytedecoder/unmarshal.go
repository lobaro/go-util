package bytedecoder

import (
	"encoding/json"
	"fmt"
	"github.com/lobaro/go-util/cborutil"
	"github.com/lobaro/go-util/jsonutil"
	"github.com/lobaro/go-util/reflectutil"
)

func UnmarshalCborOrJsonMap(v interface{}, bytes []byte) error {
	reflectutil.MustBePointer(v)
	if len(bytes) == 0 {
		return fmt.Errorf("empty body")
	}
	if cborutil.IsMap(bytes) {
		return cborutil.UnmarshalMap(v, bytes)
	}
	if jsonutil.IsMap(bytes) {
		return json.Unmarshal(bytes, v)
	}
	return fmt.Errorf("unknown payload format")
}
