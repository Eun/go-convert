package convert

import (
	"errors"
	"reflect"
	"strconv"
)

func (conv *Converter) convertToUint8(src, _ *convertValue) (reflect.Value, error) {
	if src.IsNil() {
		return reflect.Value{}, errors.New("source cannot be nil")
	}
	switch src.Base.Kind() {
	case reflect.Uint8:
		return src.Base, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(uint8(src.Base.Int())), nil
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(uint8(src.Base.Uint())), nil
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(uint8(src.Base.Float())), nil
	case reflect.Bool:
		if src.Base.Bool() {
			return reflect.ValueOf(uint8(1)), nil
		}
		return reflect.ValueOf(uint8(0)), nil
	case reflect.String:
		n, err := strconv.ParseUint(src.Base.String(), 0, 8)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(uint8(n)), nil
	}
	return reflect.Value{}, nil
}
