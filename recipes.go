package convert

import (
	"reflect"
)

type stdRecipes struct{}

func getStdRecipes() []Recipe {
	var r stdRecipes
	var s string
	return []Recipe{
		// bool
		MustMakeRecipe(r.intToBool),
		MustMakeRecipe(r.int8ToBool),
		MustMakeRecipe(r.int16ToBool),
		MustMakeRecipe(r.int32ToBool),
		MustMakeRecipe(r.int64ToBool),
		MustMakeRecipe(r.uintToBool),
		MustMakeRecipe(r.uint8ToBool),
		MustMakeRecipe(r.uint16ToBool),
		MustMakeRecipe(r.uint32ToBool),
		MustMakeRecipe(r.uint64ToBool),
		MustMakeRecipe(r.boolToBool),
		MustMakeRecipe(r.float32ToBool),
		MustMakeRecipe(r.float64ToBool),
		MustMakeRecipe(r.stringToBool),

		// float32
		MustMakeRecipe(r.intToFloat32),
		MustMakeRecipe(r.int8ToFloat32),
		MustMakeRecipe(r.int16ToFloat32),
		MustMakeRecipe(r.int32ToFloat32),
		MustMakeRecipe(r.int64ToFloat32),
		MustMakeRecipe(r.uintToFloat32),
		MustMakeRecipe(r.uint8ToFloat32),
		MustMakeRecipe(r.uint16ToFloat32),
		MustMakeRecipe(r.uint32ToFloat32),
		MustMakeRecipe(r.uint64ToFloat32),
		MustMakeRecipe(r.boolToFloat32),
		MustMakeRecipe(r.float32ToFloat32),
		MustMakeRecipe(r.float64ToFloat32),
		MustMakeRecipe(r.stringToFloat32),
		MustMakeRecipe(r.timeToFloat32),

		// float64
		MustMakeRecipe(r.intToFloat64),
		MustMakeRecipe(r.int8ToFloat64),
		MustMakeRecipe(r.int16ToFloat64),
		MustMakeRecipe(r.int32ToFloat64),
		MustMakeRecipe(r.int64ToFloat64),
		MustMakeRecipe(r.uintToFloat64),
		MustMakeRecipe(r.uint8ToFloat64),
		MustMakeRecipe(r.uint16ToFloat64),
		MustMakeRecipe(r.uint32ToFloat64),
		MustMakeRecipe(r.uint64ToFloat64),
		MustMakeRecipe(r.boolToFloat64),
		MustMakeRecipe(r.float32ToFloat64),
		MustMakeRecipe(r.float64ToFloat64),
		MustMakeRecipe(r.stringToFloat64),
		MustMakeRecipe(r.timeToFloat64),

		// int
		MustMakeRecipe(r.intToInt),
		MustMakeRecipe(r.int8ToInt),
		MustMakeRecipe(r.int16ToInt),
		MustMakeRecipe(r.int32ToInt),
		MustMakeRecipe(r.int64ToInt),
		MustMakeRecipe(r.uintToInt),
		MustMakeRecipe(r.uint8ToInt),
		MustMakeRecipe(r.uint16ToInt),
		MustMakeRecipe(r.uint32ToInt),
		MustMakeRecipe(r.uint64ToInt),
		MustMakeRecipe(r.boolToInt),
		MustMakeRecipe(r.float32ToInt),
		MustMakeRecipe(r.float64ToInt),
		MustMakeRecipe(r.stringToInt),
		MustMakeRecipe(r.timeToInt),

		// int8
		MustMakeRecipe(r.intToInt8),
		MustMakeRecipe(r.int8ToInt8),
		MustMakeRecipe(r.int16ToInt8),
		MustMakeRecipe(r.int32ToInt8),
		MustMakeRecipe(r.int64ToInt8),
		MustMakeRecipe(r.uintToInt8),
		MustMakeRecipe(r.uint8ToInt8),
		MustMakeRecipe(r.uint16ToInt8),
		MustMakeRecipe(r.uint32ToInt8),
		MustMakeRecipe(r.uint64ToInt8),
		MustMakeRecipe(r.boolToInt8),
		MustMakeRecipe(r.float32ToInt8),
		MustMakeRecipe(r.float64ToInt8),
		MustMakeRecipe(r.stringToInt8),
		MustMakeRecipe(r.timeToInt8),

		// int16
		MustMakeRecipe(r.intToInt16),
		MustMakeRecipe(r.int8ToInt16),
		MustMakeRecipe(r.int16ToInt16),
		MustMakeRecipe(r.int32ToInt16),
		MustMakeRecipe(r.int64ToInt16),
		MustMakeRecipe(r.uintToInt16),
		MustMakeRecipe(r.uint8ToInt16),
		MustMakeRecipe(r.uint16ToInt16),
		MustMakeRecipe(r.uint32ToInt16),
		MustMakeRecipe(r.uint64ToInt16),
		MustMakeRecipe(r.boolToInt16),
		MustMakeRecipe(r.float32ToInt16),
		MustMakeRecipe(r.float64ToInt16),
		MustMakeRecipe(r.stringToInt16),
		MustMakeRecipe(r.timeToInt16),

		// int32
		MustMakeRecipe(r.intToInt32),
		MustMakeRecipe(r.int8ToInt32),
		MustMakeRecipe(r.int16ToInt32),
		MustMakeRecipe(r.int32ToInt32),
		MustMakeRecipe(r.int64ToInt32),
		MustMakeRecipe(r.uintToInt32),
		MustMakeRecipe(r.uint8ToInt32),
		MustMakeRecipe(r.uint16ToInt32),
		MustMakeRecipe(r.uint32ToInt32),
		MustMakeRecipe(r.uint64ToInt32),
		MustMakeRecipe(r.boolToInt32),
		MustMakeRecipe(r.float32ToInt32),
		MustMakeRecipe(r.float64ToInt32),
		MustMakeRecipe(r.stringToInt32),
		MustMakeRecipe(r.timeToInt32),

		// int64
		MustMakeRecipe(r.intToInt64),
		MustMakeRecipe(r.int8ToInt64),
		MustMakeRecipe(r.int16ToInt64),
		MustMakeRecipe(r.int32ToInt64),
		MustMakeRecipe(r.int64ToInt64),
		MustMakeRecipe(r.uintToInt64),
		MustMakeRecipe(r.uint8ToInt64),
		MustMakeRecipe(r.uint16ToInt64),
		MustMakeRecipe(r.uint32ToInt64),
		MustMakeRecipe(r.uint64ToInt64),
		MustMakeRecipe(r.boolToInt64),
		MustMakeRecipe(r.float32ToInt64),
		MustMakeRecipe(r.float64ToInt64),
		MustMakeRecipe(r.stringToInt64),
		MustMakeRecipe(r.timeToInt64),

		// map
		{
			From: MapType,
			To:   MapType,
			Func: r.mapToMap,
		},
		{
			From: StructType,
			To:   MapType,
			Func: r.structToMap,
		},

		// string
		MustMakeRecipe(r.intToString),
		MustMakeRecipe(r.int8ToString),
		MustMakeRecipe(r.int16ToString),
		MustMakeRecipe(r.int32ToString),
		MustMakeRecipe(r.int64ToString),
		MustMakeRecipe(r.uintToString),
		MustMakeRecipe(r.uint8ToString),
		MustMakeRecipe(r.uint16ToString),
		MustMakeRecipe(r.uint32ToString),
		MustMakeRecipe(r.uint64ToString),
		MustMakeRecipe(r.boolToString),
		MustMakeRecipe(r.float32ToString),
		MustMakeRecipe(r.float64ToString),
		MustMakeRecipe(r.stringToString),
		MustMakeRecipe(r.intSliceToString),
		MustMakeRecipe(r.int8SliceToString),
		MustMakeRecipe(r.int16SliceToString),
		MustMakeRecipe(r.int32SliceToString),
		MustMakeRecipe(r.int64SliceToString),
		MustMakeRecipe(r.uintSliceToString),
		MustMakeRecipe(r.uint8SliceToString),
		MustMakeRecipe(r.uint16SliceToString),
		MustMakeRecipe(r.uint32SliceToString),
		MustMakeRecipe(r.uint64SliceToString),
		MustMakeRecipe(r.timeToString),

		{
			From: StructType,
			To:   reflect.TypeOf(&s),
			Func: r.structToString,
		},

		// slice
		{
			From: reflect.TypeOf(""),
			To:   SliceType,
			Func: r.stringToSlice,
		},
		{
			From: SliceType,
			To:   SliceType,
			Func: r.sliceToSlice,
		},
		{
			From: NilType,
			To:   SliceType,
			Func: r.nilToSlice,
		},

		// struct
		{
			From: MapType,
			To:   StructType,
			Func: r.mapToStruct,
		},
		{
			From: StructType,
			To:   StructType,
			Func: r.structToStruct,
		},

		// time
		MustMakeRecipe(r.intToTime),
		MustMakeRecipe(r.int8ToTime),
		MustMakeRecipe(r.int16ToTime),
		MustMakeRecipe(r.int32ToTime),
		MustMakeRecipe(r.int64ToTime),
		MustMakeRecipe(r.uintToTime),
		MustMakeRecipe(r.uint8ToTime),
		MustMakeRecipe(r.uint16ToTime),
		MustMakeRecipe(r.uint32ToTime),
		MustMakeRecipe(r.uint64ToTime),
		MustMakeRecipe(r.float32ToTime),
		MustMakeRecipe(r.float64ToTime),
		MustMakeRecipe(r.stringToTime),

		// uint
		MustMakeRecipe(r.intToUint),
		MustMakeRecipe(r.int8ToUint),
		MustMakeRecipe(r.int16ToUint),
		MustMakeRecipe(r.int32ToUint),
		MustMakeRecipe(r.int64ToUint),
		MustMakeRecipe(r.uintToUint),
		MustMakeRecipe(r.uint8ToUint),
		MustMakeRecipe(r.uint16ToUint),
		MustMakeRecipe(r.uint32ToUint),
		MustMakeRecipe(r.uint64ToUint),
		MustMakeRecipe(r.boolToUint),
		MustMakeRecipe(r.float32ToUint),
		MustMakeRecipe(r.float64ToUint),
		MustMakeRecipe(r.stringToUint),
		MustMakeRecipe(r.timeToUint),

		// uint8
		MustMakeRecipe(r.intToUint8),
		MustMakeRecipe(r.int8ToUint8),
		MustMakeRecipe(r.int16ToUint8),
		MustMakeRecipe(r.int32ToUint8),
		MustMakeRecipe(r.int64ToUint8),
		MustMakeRecipe(r.uintToUint8),
		MustMakeRecipe(r.uint8ToUint8),
		MustMakeRecipe(r.uint16ToUint8),
		MustMakeRecipe(r.uint32ToUint8),
		MustMakeRecipe(r.uint64ToUint8),
		MustMakeRecipe(r.boolToUint8),
		MustMakeRecipe(r.float32ToUint8),
		MustMakeRecipe(r.float64ToUint8),
		MustMakeRecipe(r.stringToUint8),
		MustMakeRecipe(r.timeToUint8),

		// uint16
		MustMakeRecipe(r.intToUint16),
		MustMakeRecipe(r.int8ToUint16),
		MustMakeRecipe(r.int16ToUint16),
		MustMakeRecipe(r.int32ToUint16),
		MustMakeRecipe(r.int64ToUint16),
		MustMakeRecipe(r.uintToUint16),
		MustMakeRecipe(r.uint8ToUint16),
		MustMakeRecipe(r.uint16ToUint16),
		MustMakeRecipe(r.uint32ToUint16),
		MustMakeRecipe(r.uint64ToUint16),
		MustMakeRecipe(r.boolToUint16),
		MustMakeRecipe(r.float32ToUint16),
		MustMakeRecipe(r.float64ToUint16),
		MustMakeRecipe(r.stringToUint16),
		MustMakeRecipe(r.timeToUint16),

		// uint32
		MustMakeRecipe(r.intToUint32),
		MustMakeRecipe(r.int8ToUint32),
		MustMakeRecipe(r.int16ToUint32),
		MustMakeRecipe(r.int32ToUint32),
		MustMakeRecipe(r.int64ToUint32),
		MustMakeRecipe(r.uintToUint32),
		MustMakeRecipe(r.uint8ToUint32),
		MustMakeRecipe(r.uint16ToUint32),
		MustMakeRecipe(r.uint32ToUint32),
		MustMakeRecipe(r.uint64ToUint32),
		MustMakeRecipe(r.boolToUint32),
		MustMakeRecipe(r.float32ToUint32),
		MustMakeRecipe(r.float64ToUint32),
		MustMakeRecipe(r.stringToUint32),
		MustMakeRecipe(r.timeToUint32),

		// uint64
		MustMakeRecipe(r.intToUint64),
		MustMakeRecipe(r.int8ToUint64),
		MustMakeRecipe(r.int16ToUint64),
		MustMakeRecipe(r.int32ToUint64),
		MustMakeRecipe(r.int64ToUint64),
		MustMakeRecipe(r.uintToUint64),
		MustMakeRecipe(r.uint8ToUint64),
		MustMakeRecipe(r.uint16ToUint64),
		MustMakeRecipe(r.uint32ToUint64),
		MustMakeRecipe(r.uint64ToUint64),
		MustMakeRecipe(r.boolToUint64),
		MustMakeRecipe(r.float32ToUint64),
		MustMakeRecipe(r.float64ToUint64),
		MustMakeRecipe(r.stringToUint64),
		MustMakeRecipe(r.timeToUint64),
	}
}
