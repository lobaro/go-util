package cbor

import (
	"io"

	"github.com/ugorji/go/codec"
)

func NewDecoder(r io.Reader) *codec.Decoder {
	return codec.NewDecoder(r, new(codec.CborHandle))
}

func NewEncoder(w io.Writer) *codec.Encoder {
	return codec.NewEncoder(w, new(codec.CborHandle))
}

func NewDecoderBytes(data []byte) *codec.Decoder {
	return codec.NewDecoderBytes(data, new(codec.CborHandle))
}

func NewEncoderBytes(out *[]byte) *codec.Encoder {
	return codec.NewEncoderBytes(out, new(codec.CborHandle))
}

func Marshal(v interface{}) ([]byte, error) {
	b := make([]byte, 0)
	enc := NewEncoderBytes(&b)
	err := enc.Encode(v)
	return b, err
}

func Unmarshal(data []byte, v interface{}) error {
	dec := NewDecoderBytes(data)
	err := dec.Decode(v)

	return err
}
