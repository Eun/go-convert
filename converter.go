package convert

import "reflect"

type NilValue struct{}
type MapValue struct{}
type StructValue struct{}
type SliceValue struct{}

// NilType can be used to specify a recipe with the source/destination with a nil value
var NilType = reflect.TypeOf((*NilValue)(nil)).Elem()

// MapType can be used to specify a recipe with the source/destination with a map type
var MapType = reflect.TypeOf((*MapValue)(nil)).Elem()

// StructType can be used to specify a recipe with the source/destination with a struct type
var StructType = reflect.TypeOf((*StructValue)(nil)).Elem()

// SliceType can be used to specify a recipe with the source/destination with a slice type
var SliceType = reflect.TypeOf((*SliceValue)(nil)).Elem()

// Converter is the instance that will be used to convert values
type Converter interface {
	Options() *Options
	Convert(src, dst interface{}, options ...Options) error
	MustConvert(src, dst interface{}, options ...Options)
	ConvertReflectValue(src, dst reflect.Value, options ...Options) error
	MustConvertReflectValue(src, dst reflect.Value, options ...Options)
}
