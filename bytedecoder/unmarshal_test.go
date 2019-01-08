package bytedecoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyStruct struct {
	User PersonStruct `json:"user"`
	Id   int64        `json:"id"`
}
type PersonStruct struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint8  `json:"age"`
}

var expectedStruct = MyStruct{
	PersonStruct{"Max", "Mustermann-Müller", 18},
	-123456789,
}

const jsonStruct = `{"user":{"first_name":"Max","last_name":"Mustermann-Müller","age":18},"id":-123456789}`
const cborStruct = "\xa2duser\xa3jfirst_namecMaxilast_namerMustermann-M\xc3\xbcllercage\x12bid:\x07[\xcd\x14"
const jsonArray = `[1,2,3,null,"abcdefg","abcdefg",1234567890123456,"this is a somewhat longer string that needs a length byte"]`
const cborArray = "\x88\x01\x02\x03\xf6gabcdefgGabcdefg\x1b\x00\x04b\xd5<\x8a\xba\xc0x9this is a somewhat longer string that needs a length byte"

func TestUnmarshalJsonMap(t *testing.T) {
	ms := MyStruct{}
	err := UnmarshalCborOrJson([]byte(jsonStruct), &ms)
	assert.NoError(t, err)
	assert.Equal(t, expectedStruct, ms)
}

func TestUnmarshalCborMap(t *testing.T) {
	ms := MyStruct{}
	err := UnmarshalCborOrJson([]byte(cborStruct), &ms)
	assert.NoError(t, err)
	assert.Equal(t, expectedStruct, ms)
}

func TestUnmarshalJsonArray(t *testing.T) {
	a := [8]interface{}{}
	err := UnmarshalCborOrJson([]byte(jsonArray), &a)
	assert.NoError(t, err)
	assert.Equal(t, 8, len(a))
	assert.EqualValues(t, 1, a[0])
	assert.EqualValues(t, 2, a[1])
	assert.EqualValues(t, 3, a[2])
	assert.EqualValues(t, nil, a[3])
	assert.EqualValues(t, "abcdefg", a[4])
	assert.EqualValues(t, []byte("abcdefg"), a[5])
	assert.EqualValues(t, 1234567890123456, a[6])
	assert.EqualValues(t, "this is a somewhat longer string that needs a length byte", a[7])
}

func TestUnmarshalCborArray(t *testing.T) {
	a := [8]interface{}{}
	err := UnmarshalCborOrJson([]byte(cborArray), &a)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, 8, len(a))
	assert.EqualValues(t, 1, a[0])
	assert.EqualValues(t, 2, a[1])
	assert.EqualValues(t, 3, a[2])
	assert.EqualValues(t, nil, a[3])
	assert.EqualValues(t, "abcdefg", a[4])
	assert.EqualValues(t, []byte("abcdefg"), a[5])
	assert.EqualValues(t, 1234567890123456, a[6])
	assert.EqualValues(t, "this is a somewhat longer string that needs a length byte", a[7])
}
