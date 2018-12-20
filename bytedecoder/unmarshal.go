package bytedecoder

import (
	"encoding/json"
	"fmt"
	"github.com/lobaro/go-util/cbor"
	"github.com/lobaro/go-util/jsonutil"
	"github.com/lobaro/go-util/reflectutil"
)

// convenience function to extract interface from bytes from either cbor or json
func UnmarshalCborOrJson(bytes []byte, v interface{}) error {
	reflectutil.MustBePointer(v)
	if len(bytes) == 0 {
		return fmt.Errorf("empty body")
	}
	if cbor.IsMap(bytes) || cbor.IsArray(bytes) {
		return cbor.Unmarshal(bytes, v)
	}
	if jsonutil.IsMap(bytes) || jsonutil.IsArray(bytes) {
		return json.Unmarshal(bytes, v)
	}
	return fmt.Errorf("unknown payload format")
}
