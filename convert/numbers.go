package convert

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

var float64Type = reflect.TypeOf(float64(0))
var intType = reflect.TypeOf(int(0))
var stringType = reflect.TypeOf("")

func ToFloat(val interface{}) (float64, error) {
	switch i := val.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	default:
		v := reflect.ValueOf(val)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(float64Type) {
			fv := v.Convert(float64Type)
			return fv.Float(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			return strconv.ParseFloat(s, 64)
		} else {
			return math.NaN(), fmt.Errorf("can not convert %v (%T) to float64", v, v)
		}
	}
}

func ToInt(val interface{}) (int, error) {
	switch i := val.(type) {
	case float64:
		return int(i), nil
	case float32:
		return int(i), nil
	case int64:
		return int(i), nil
	case int32:
		return int(i), nil
	case int:
		return int(i), nil
	case uint64:
		return int(i), nil
	case uint32:
		return int(i), nil
	case uint:
		return int(i), nil
	case string:
		return strconv.Atoi(i)
	default:
		v := reflect.ValueOf(val)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(intType) {
			fv := v.Convert(intType)
			return int(fv.Int()), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			return strconv.Atoi(s)
		} else {
			return 0, fmt.Errorf("can not convert %v (%T) to int", v, v)
		}
	}
}
