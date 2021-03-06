package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithInt8Func struct{}

func (SomeStructWithInt8Func) Int8() int8 {
	return 8
}

type SomeStructWithInt8FuncPtr struct{}

func (*SomeStructWithInt8FuncPtr) Int8() int8 {
	return 8
}

type SomeStructWithInt8WithErrFunc struct{}

func (SomeStructWithInt8WithErrFunc) Int8() (int8, error) {
	return 8, nil
}

type SomeStructWithInt8WithErrFuncPtr struct{}

func (*SomeStructWithInt8WithErrFuncPtr) Int8() (int8, error) {
	return 8, nil
}

func TestInt8(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, int8(0), int8(0), "", nil},
		// string
		{"6", int8(0), int8(6), "", nil},
		{"", int8(0), int8(0), "", nil},
		{"Hello World", int8(0), int8(0), `unable to convert string to int8: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, int8(0), int8(1), "", nil},
		{false, int8(0), int8(0), "", nil},
		// int
		{6, int8(0), int8(6), "", nil},
		// int8
		{int8(6), int8(0), int8(6), "", nil},
		// int16
		{int16(6), int8(0), int8(6), "", nil},
		// int32
		{int32(6), int8(0), int8(6), "", nil},
		// int64
		{int64(6), int8(0), int8(6), "", nil},
		// uint
		{uint(6), int8(0), int8(6), "", nil},
		// uint8
		{uint8(6), int8(0), int8(6), "", nil},
		// uint16
		{uint16(6), int8(0), int8(6), "", nil},
		// uint32
		{uint32(6), int8(0), int8(6), "", nil},
		// uint64
		{uint64(6), int8(0), int8(6), "", nil},
		// float32
		{float32(6), int8(0), int8(6), "", nil},
		// float64
		{float64(6), int8(0), int8(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, int8(0), int8(0), "unable to convert []int to int8: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, int8(0), int8(0), "unable to convert []uint8 to int8: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, int8(0), int8(0), "unable to convert []int32 to int8: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, int8(0), int8(0), "unable to convert []string to int8: no recipe", nil},
		// struct
		{struct{}{}, int8(0), int8(0), "unable to convert struct {} to int8: struct {} has no Int8() function", nil},
		// time
		{time.Unix(10, 10), int8(10), int8(10), "", nil},

		{SomeStructWithIntFunc{}, int8(0), int8(8), "", nil},
		{&SomeStructWithIntFunc{}, int8(0), int8(8), "", nil},

		{SomeStructWithIntFuncPtr{}, int8(0), int8(8), "", nil},
		{&SomeStructWithIntFuncPtr{}, int8(0), int8(8), "", nil},

		{SomeStructWithIntWithErrFunc{}, int8(0), int8(8), "", nil},
		{&SomeStructWithIntWithErrFunc{}, int8(0), int8(8), "", nil},

		{SomeStructWithIntWithErrFuncPtr{}, int8(0), int8(8), "", nil},
		{&SomeStructWithIntWithErrFuncPtr{}, int8(0), int8(8), "", nil},

		{SomeStructWithInt8Func{}, int8(0), int8(8), "", nil},
		{&SomeStructWithInt8Func{}, int8(0), int8(8), "", nil},

		{SomeStructWithInt8FuncPtr{}, int8(0), int8(8), "", nil},
		{&SomeStructWithInt8FuncPtr{}, int8(0), int8(8), "", nil},

		{SomeStructWithInt8WithErrFunc{}, int8(0), int8(8), "", nil},
		{&SomeStructWithInt8WithErrFunc{}, int8(0), int8(8), "", nil},

		{SomeStructWithInt8WithErrFuncPtr{}, int8(0), int8(8), "", nil},
		{&SomeStructWithInt8WithErrFuncPtr{}, int8(0), int8(8), "", nil},

		{SomeStructWithInt16Func{}, int8(0), int8(16), "", nil},
		{&SomeStructWithInt16Func{}, int8(0), int8(16), "", nil},

		{SomeStructWithInt16FuncPtr{}, int8(0), int8(16), "", nil},
		{&SomeStructWithInt16FuncPtr{}, int8(0), int8(16), "", nil},

		{SomeStructWithInt16WithErrFunc{}, int8(0), int8(16), "", nil},
		{&SomeStructWithInt16WithErrFunc{}, int8(0), int8(16), "", nil},

		{SomeStructWithInt16WithErrFuncPtr{}, int8(0), int8(16), "", nil},
		{&SomeStructWithInt16WithErrFuncPtr{}, int8(0), int8(16), "", nil},

		{SomeStructWithInt32Func{}, int8(0), int8(32), "", nil},
		{&SomeStructWithInt32Func{}, int8(0), int8(32), "", nil},

		{SomeStructWithInt32FuncPtr{}, int8(0), int8(32), "", nil},
		{&SomeStructWithInt32FuncPtr{}, int8(0), int8(32), "", nil},

		{SomeStructWithInt32WithErrFunc{}, int8(0), int8(32), "", nil},
		{&SomeStructWithInt32WithErrFunc{}, int8(0), int8(32), "", nil},

		{SomeStructWithInt32WithErrFuncPtr{}, int8(0), int8(32), "", nil},
		{&SomeStructWithInt32WithErrFuncPtr{}, int8(0), int8(32), "", nil},

		{SomeStructWithInt64Func{}, int8(0), int8(64), "", nil},
		{&SomeStructWithInt64Func{}, int8(0), int8(64), "", nil},

		{SomeStructWithInt64FuncPtr{}, int8(0), int8(64), "", nil},
		{&SomeStructWithInt64FuncPtr{}, int8(0), int8(64), "", nil},

		{SomeStructWithInt64WithErrFunc{}, int8(0), int8(64), "", nil},
		{&SomeStructWithInt64WithErrFunc{}, int8(0), int8(64), "", nil},

		{SomeStructWithInt64WithErrFuncPtr{}, int8(0), int8(64), "", nil},
		{&SomeStructWithInt64WithErrFuncPtr{}, int8(0), int8(64), "", nil},

		{SomeStructWithUintFunc{}, int8(0), int8(16), "", nil},
		{&SomeStructWithUintFunc{}, int8(0), int8(16), "", nil},

		{SomeStructWithUintFuncPtr{}, int8(0), int8(16), "", nil},
		{&SomeStructWithUintFuncPtr{}, int8(0), int8(16), "", nil},

		{SomeStructWithUintWithErrFunc{}, int8(0), int8(16), "", nil},
		{&SomeStructWithUintWithErrFunc{}, int8(0), int8(16), "", nil},

		{SomeStructWithUintWithErrFuncPtr{}, int8(0), int8(16), "", nil},
		{&SomeStructWithUintWithErrFuncPtr{}, int8(0), int8(16), "", nil},

		{SomeStructWithUint8Func{}, int8(0), int8(8), "", nil},
		{&SomeStructWithUint8Func{}, int8(0), int8(8), "", nil},

		{SomeStructWithUint8FuncPtr{}, int8(0), int8(8), "", nil},
		{&SomeStructWithUint8FuncPtr{}, int8(0), int8(8), "", nil},

		{SomeStructWithUint8WithErrFunc{}, int8(0), int8(8), "", nil},
		{&SomeStructWithUint8WithErrFunc{}, int8(0), int8(8), "", nil},

		{SomeStructWithUint8WithErrFuncPtr{}, int8(0), int8(8), "", nil},
		{&SomeStructWithUint8WithErrFuncPtr{}, int8(0), int8(8), "", nil},

		{SomeStructWithUint16Func{}, int8(0), int8(16), "", nil},
		{&SomeStructWithUint16Func{}, int8(0), int8(16), "", nil},

		{SomeStructWithUint16FuncPtr{}, int8(0), int8(16), "", nil},
		{&SomeStructWithUint16FuncPtr{}, int8(0), int8(16), "", nil},

		{SomeStructWithUint16WithErrFunc{}, int8(0), int8(16), "", nil},
		{&SomeStructWithUint16WithErrFunc{}, int8(0), int8(16), "", nil},

		{SomeStructWithUint16WithErrFuncPtr{}, int8(0), int8(16), "", nil},
		{&SomeStructWithUint16WithErrFuncPtr{}, int8(0), int8(16), "", nil},

		{SomeStructWithUint32Func{}, int8(0), int8(32), "", nil},
		{&SomeStructWithUint32Func{}, int8(0), int8(32), "", nil},

		{SomeStructWithUint32FuncPtr{}, int8(0), int8(32), "", nil},
		{&SomeStructWithUint32FuncPtr{}, int8(0), int8(32), "", nil},

		{SomeStructWithUint32WithErrFunc{}, int8(0), int8(32), "", nil},
		{&SomeStructWithUint32WithErrFunc{}, int8(0), int8(32), "", nil},

		{SomeStructWithUint32WithErrFuncPtr{}, int8(0), int8(32), "", nil},
		{&SomeStructWithUint32WithErrFuncPtr{}, int8(0), int8(32), "", nil},

		{SomeStructWithUint64Func{}, int8(0), int8(64), "", nil},
		{&SomeStructWithUint64Func{}, int8(0), int8(64), "", nil},

		{SomeStructWithUint64FuncPtr{}, int8(0), int8(64), "", nil},
		{&SomeStructWithUint64FuncPtr{}, int8(0), int8(64), "", nil},

		{SomeStructWithUint64WithErrFunc{}, int8(0), int8(64), "", nil},
		{&SomeStructWithUint64WithErrFunc{}, int8(0), int8(64), "", nil},

		{SomeStructWithUint64WithErrFuncPtr{}, int8(0), int8(64), "", nil},
		{&SomeStructWithUint64WithErrFuncPtr{}, int8(0), int8(64), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
