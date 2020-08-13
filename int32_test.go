package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithInt32Func struct{}

func (SomeStructWithInt32Func) Int32() int32 {
	return 32
}

type SomeStructWithInt32FuncPtr struct{}

func (*SomeStructWithInt32FuncPtr) Int32() int32 {
	return 32
}

type SomeStructWithInt32WithErrFunc struct{}

func (SomeStructWithInt32WithErrFunc) Int32() (int32, error) {
	return 32, nil
}

type SomeStructWithInt32WithErrFuncPtr struct{}

func (*SomeStructWithInt32WithErrFuncPtr) Int32() (int32, error) {
	return 32, nil
}

func TestInt32(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, int32(0), int32(0), "", nil},
		// string
		{"6", int32(0), int32(6), "", nil},
		{"", int32(0), int32(0), "", nil},
		{"Hello World", int32(0), int32(0), `unable to convert string to int32: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, int32(0), int32(1), "", nil},
		{false, int32(0), int32(0), "", nil},
		// int
		{6, int32(0), int32(6), "", nil},
		// int8
		{int8(6), int32(0), int32(6), "", nil},
		// int16
		{int16(6), int32(0), int32(6), "", nil},
		// int32
		{int32(6), int32(0), int32(6), "", nil},
		// int64
		{int64(6), int32(0), int32(6), "", nil},
		// uint
		{uint(6), int32(0), int32(6), "", nil},
		// uint8
		{uint8(6), int32(0), int32(6), "", nil},
		// uint16
		{uint16(6), int32(0), int32(6), "", nil},
		// uint32
		{uint32(6), int32(0), int32(6), "", nil},
		// uint64
		{uint64(6), int32(0), int32(6), "", nil},
		// float32
		{float32(6), int32(0), int32(6), "", nil},
		// float64
		{float64(6), int32(0), int32(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, int32(0), int32(0), "unable to convert []int to int32: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, int32(0), int32(0), "unable to convert []uint8 to int32: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, int32(0), int32(0), "unable to convert []int32 to int32: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, int32(0), int32(0), "unable to convert []string to int32: no recipe", nil},
		// struct
		{struct{}{}, int32(0), int32(0), "unable to convert struct {} to int32: struct {} has no Int32() function", nil},
		// time
		{time.Unix(10, 10), int32(10), int32(10), "", nil},

		{SomeStructWithIntFunc{}, int32(0), int32(8), "", nil},
		{&SomeStructWithIntFunc{}, int32(0), int32(8), "", nil},

		{SomeStructWithIntFuncPtr{}, int32(0), int32(8), "", nil},
		{&SomeStructWithIntFuncPtr{}, int32(0), int32(8), "", nil},

		{SomeStructWithIntWithErrFunc{}, int32(0), int32(8), "", nil},
		{&SomeStructWithIntWithErrFunc{}, int32(0), int32(8), "", nil},

		{SomeStructWithIntWithErrFuncPtr{}, int32(0), int32(8), "", nil},
		{&SomeStructWithIntWithErrFuncPtr{}, int32(0), int32(8), "", nil},

		{SomeStructWithInt8Func{}, int32(0), int32(8), "", nil},
		{&SomeStructWithInt8Func{}, int32(0), int32(8), "", nil},

		{SomeStructWithInt8FuncPtr{}, int32(0), int32(8), "", nil},
		{&SomeStructWithInt8FuncPtr{}, int32(0), int32(8), "", nil},

		{SomeStructWithInt8WithErrFunc{}, int32(0), int32(8), "", nil},
		{&SomeStructWithInt8WithErrFunc{}, int32(0), int32(8), "", nil},

		{SomeStructWithInt8WithErrFuncPtr{}, int32(0), int32(8), "", nil},
		{&SomeStructWithInt8WithErrFuncPtr{}, int32(0), int32(8), "", nil},

		{SomeStructWithInt16Func{}, int32(0), int32(16), "", nil},
		{&SomeStructWithInt16Func{}, int32(0), int32(16), "", nil},

		{SomeStructWithInt16FuncPtr{}, int32(0), int32(16), "", nil},
		{&SomeStructWithInt16FuncPtr{}, int32(0), int32(16), "", nil},

		{SomeStructWithInt16WithErrFunc{}, int32(0), int32(16), "", nil},
		{&SomeStructWithInt16WithErrFunc{}, int32(0), int32(16), "", nil},

		{SomeStructWithInt16WithErrFuncPtr{}, int32(0), int32(16), "", nil},
		{&SomeStructWithInt16WithErrFuncPtr{}, int32(0), int32(16), "", nil},

		{SomeStructWithInt32Func{}, int32(0), int32(32), "", nil},
		{&SomeStructWithInt32Func{}, int32(0), int32(32), "", nil},

		{SomeStructWithInt32FuncPtr{}, int32(0), int32(32), "", nil},
		{&SomeStructWithInt32FuncPtr{}, int32(0), int32(32), "", nil},

		{SomeStructWithInt32WithErrFunc{}, int32(0), int32(32), "", nil},
		{&SomeStructWithInt32WithErrFunc{}, int32(0), int32(32), "", nil},

		{SomeStructWithInt32WithErrFuncPtr{}, int32(0), int32(32), "", nil},
		{&SomeStructWithInt32WithErrFuncPtr{}, int32(0), int32(32), "", nil},

		{SomeStructWithInt64Func{}, int32(0), int32(64), "", nil},
		{&SomeStructWithInt64Func{}, int32(0), int32(64), "", nil},

		{SomeStructWithInt64FuncPtr{}, int32(0), int32(64), "", nil},
		{&SomeStructWithInt64FuncPtr{}, int32(0), int32(64), "", nil},

		{SomeStructWithInt64WithErrFunc{}, int32(0), int32(64), "", nil},
		{&SomeStructWithInt64WithErrFunc{}, int32(0), int32(64), "", nil},

		{SomeStructWithInt64WithErrFuncPtr{}, int32(0), int32(64), "", nil},
		{&SomeStructWithInt64WithErrFuncPtr{}, int32(0), int32(64), "", nil},

		{SomeStructWithUintFunc{}, int32(0), int32(16), "", nil},
		{&SomeStructWithUintFunc{}, int32(0), int32(16), "", nil},

		{SomeStructWithUintFuncPtr{}, int32(0), int32(16), "", nil},
		{&SomeStructWithUintFuncPtr{}, int32(0), int32(16), "", nil},

		{SomeStructWithUintWithErrFunc{}, int32(0), int32(16), "", nil},
		{&SomeStructWithUintWithErrFunc{}, int32(0), int32(16), "", nil},

		{SomeStructWithUintWithErrFuncPtr{}, int32(0), int32(16), "", nil},
		{&SomeStructWithUintWithErrFuncPtr{}, int32(0), int32(16), "", nil},

		{SomeStructWithUint8Func{}, int32(0), int32(8), "", nil},
		{&SomeStructWithUint8Func{}, int32(0), int32(8), "", nil},

		{SomeStructWithUint8FuncPtr{}, int32(0), int32(8), "", nil},
		{&SomeStructWithUint8FuncPtr{}, int32(0), int32(8), "", nil},

		{SomeStructWithUint8WithErrFunc{}, int32(0), int32(8), "", nil},
		{&SomeStructWithUint8WithErrFunc{}, int32(0), int32(8), "", nil},

		{SomeStructWithUint8WithErrFuncPtr{}, int32(0), int32(8), "", nil},
		{&SomeStructWithUint8WithErrFuncPtr{}, int32(0), int32(8), "", nil},

		{SomeStructWithUint16Func{}, int32(0), int32(16), "", nil},
		{&SomeStructWithUint16Func{}, int32(0), int32(16), "", nil},

		{SomeStructWithUint16FuncPtr{}, int32(0), int32(16), "", nil},
		{&SomeStructWithUint16FuncPtr{}, int32(0), int32(16), "", nil},

		{SomeStructWithUint16WithErrFunc{}, int32(0), int32(16), "", nil},
		{&SomeStructWithUint16WithErrFunc{}, int32(0), int32(16), "", nil},

		{SomeStructWithUint16WithErrFuncPtr{}, int32(0), int32(16), "", nil},
		{&SomeStructWithUint16WithErrFuncPtr{}, int32(0), int32(16), "", nil},

		{SomeStructWithUint32Func{}, int32(0), int32(32), "", nil},
		{&SomeStructWithUint32Func{}, int32(0), int32(32), "", nil},

		{SomeStructWithUint32FuncPtr{}, int32(0), int32(32), "", nil},
		{&SomeStructWithUint32FuncPtr{}, int32(0), int32(32), "", nil},

		{SomeStructWithUint32WithErrFunc{}, int32(0), int32(32), "", nil},
		{&SomeStructWithUint32WithErrFunc{}, int32(0), int32(32), "", nil},

		{SomeStructWithUint32WithErrFuncPtr{}, int32(0), int32(32), "", nil},
		{&SomeStructWithUint32WithErrFuncPtr{}, int32(0), int32(32), "", nil},

		{SomeStructWithUint64Func{}, int32(0), int32(64), "", nil},
		{&SomeStructWithUint64Func{}, int32(0), int32(64), "", nil},

		{SomeStructWithUint64FuncPtr{}, int32(0), int32(64), "", nil},
		{&SomeStructWithUint64FuncPtr{}, int32(0), int32(64), "", nil},

		{SomeStructWithUint64WithErrFunc{}, int32(0), int32(64), "", nil},
		{&SomeStructWithUint64WithErrFunc{}, int32(0), int32(64), "", nil},

		{SomeStructWithUint64WithErrFuncPtr{}, int32(0), int32(64), "", nil},
		{&SomeStructWithUint64WithErrFuncPtr{}, int32(0), int32(64), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
