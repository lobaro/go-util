package cborutil

import (
	"github.com/ugorji/go/codec"
	"reflect"
)

type CborType byte

const (
	CborMap   = CborType(0xA0)
	CborArray = CborType(0x80)
)

func MajorType(data byte) CborType {
	return CborType(data & 0xE0)
}

func IsMap(bytes []byte) bool {
	return len(bytes) > 1 && MajorType(bytes[0]) == CborMap
}

func UnmarshalMap(v interface{}, bytes []byte) error {
	cborHandle := new(codec.CborHandle)
	cborHandle.MapType = reflect.TypeOf(make(map[string]interface{}))
	cbor := codec.NewDecoderBytes(bytes, cborHandle)
	return cbor.Decode(v)
}
/*
   pseudo code for well form check, taken from RFC 7049
   https://tools.ietf.org/html/rfc7049#appendix-C

   well_formed (breakable = false) {
     // process initial bytes
     ib = uint(take(1));
     mt = ib >> 5;
     val = ai = ib & 0x1f;
     switch (ai) {
       case 24: val = uint(take(1)); break;
       case 25: val = uint(take(2)); break;
       case 26: val = uint(take(4)); break;
       case 27: val = uint(take(8)); break;
       case 28: case 29: case 30: fail();
       case 31:
         return well_formed_indefinite(mt, breakable);
     }
     // process content
     switch (mt) {
       // case 0, 1, 7 do not have content; just use val
       case 2: case 3: take(val); break; // bytes/UTF-8
       case 4: for (i = 0; i < val; i++) well_formed(); break;
       case 5: for (i = 0; i < val*2; i++) well_formed(); break;
       case 6: well_formed(); break;     // 1 embedded data item
     }
     return mt;                    // finite data item
   }

   well_formed_indefinite(mt, breakable) {
     switch (mt) {
       case 2: case 3:
         while ((it = well_formed(true)) != -1)
           if (it != mt)           // need finite embedded
             fail();               //    of same type
         break;
       case 4: while (well_formed(true) != -1); break;
       case 5: while (well_formed(true) != -1) well_formed(); break;
       case 7:
         if (breakable)
           return -1;              // signal break out
         else fail();              // no enclosing indefinite
       default: fail();            // wrong mt
     }
     return 0;                     // no break out
   }
 */
