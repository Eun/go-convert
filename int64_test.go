package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithInt64Func struct{}

func (SomeStructWithInt64Func) Int64() int64 {
	return 64
}

type SomeStructWithInt64FuncPtr struct{}

func (*SomeStructWithInt64FuncPtr) Int64() int64 {
	return 64
}

type SomeStructWithInt64WithErrFunc struct{}

func (SomeStructWithInt64WithErrFunc) Int64() (int64, error) {
	return 64, nil
}

type SomeStructWithInt64WithErrFuncPtr struct{}

func (*SomeStructWithInt64WithErrFuncPtr) Int64() (int64, error) {
	return 64, nil
}

func TestInt64(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, int64(0), int64(0), "", nil},
		// string
		{"6", int64(0), int64(6), "", nil},
		{"", int64(0), int64(0), "", nil},
		{"Hello World", int64(0), int64(0), `unable to convert string to int64: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, int64(0), int64(1), "", nil},
		{false, int64(0), int64(0), "", nil},
		// int
		{6, int64(0), int64(6), "", nil},
		// int8
		{int8(6), int64(0), int64(6), "", nil},
		// int16
		{int16(6), int64(0), int64(6), "", nil},
		// int32
		{int32(6), int64(0), int64(6), "", nil},
		// int64
		{int64(6), int64(0), int64(6), "", nil},
		// uint
		{uint(6), int64(0), int64(6), "", nil},
		// uint8
		{uint8(6), int64(0), int64(6), "", nil},
		// uint16
		{uint16(6), int64(0), int64(6), "", nil},
		// uint32
		{uint32(6), int64(0), int64(6), "", nil},
		// uint64
		{uint64(6), int64(0), int64(6), "", nil},
		// float32
		{float32(6), int64(0), int64(6), "", nil},
		// float64
		{float64(6), int64(0), int64(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, int64(0), int64(0), "unable to convert []int to int64: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, int64(0), int64(0), "unable to convert []uint8 to int64: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, int64(0), int64(0), "unable to convert []int32 to int64: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, int64(0), int64(0), "unable to convert []string to int64: no recipe", nil},
		// struct
		{struct{}{}, int64(0), int64(0), "unable to convert struct {} to int64: struct {} has no Int64() function", nil},
		// time
		{time.Unix(10, 10), int64(10), int64(10), "", nil},

		{SomeStructWithIntFunc{}, int64(0), int64(8), "", nil},
		{&SomeStructWithIntFunc{}, int64(0), int64(8), "", nil},

		{SomeStructWithIntFuncPtr{}, int64(0), int64(8), "", nil},
		{&SomeStructWithIntFuncPtr{}, int64(0), int64(8), "", nil},

		{SomeStructWithIntWithErrFunc{}, int64(0), int64(8), "", nil},
		{&SomeStructWithIntWithErrFunc{}, int64(0), int64(8), "", nil},

		{SomeStructWithIntWithErrFuncPtr{}, int64(0), int64(8), "", nil},
		{&SomeStructWithIntWithErrFuncPtr{}, int64(0), int64(8), "", nil},

		{SomeStructWithInt8Func{}, int64(0), int64(8), "", nil},
		{&SomeStructWithInt8Func{}, int64(0), int64(8), "", nil},

		{SomeStructWithInt8FuncPtr{}, int64(0), int64(8), "", nil},
		{&SomeStructWithInt8FuncPtr{}, int64(0), int64(8), "", nil},

		{SomeStructWithInt8WithErrFunc{}, int64(0), int64(8), "", nil},
		{&SomeStructWithInt8WithErrFunc{}, int64(0), int64(8), "", nil},

		{SomeStructWithInt8WithErrFuncPtr{}, int64(0), int64(8), "", nil},
		{&SomeStructWithInt8WithErrFuncPtr{}, int64(0), int64(8), "", nil},

		{SomeStructWithInt16Func{}, int64(0), int64(16), "", nil},
		{&SomeStructWithInt16Func{}, int64(0), int64(16), "", nil},

		{SomeStructWithInt16FuncPtr{}, int64(0), int64(16), "", nil},
		{&SomeStructWithInt16FuncPtr{}, int64(0), int64(16), "", nil},

		{SomeStructWithInt16WithErrFunc{}, int64(0), int64(16), "", nil},
		{&SomeStructWithInt16WithErrFunc{}, int64(0), int64(16), "", nil},

		{SomeStructWithInt16WithErrFuncPtr{}, int64(0), int64(16), "", nil},
		{&SomeStructWithInt16WithErrFuncPtr{}, int64(0), int64(16), "", nil},

		{SomeStructWithInt32Func{}, int64(0), int64(32), "", nil},
		{&SomeStructWithInt32Func{}, int64(0), int64(32), "", nil},

		{SomeStructWithInt32FuncPtr{}, int64(0), int64(32), "", nil},
		{&SomeStructWithInt32FuncPtr{}, int64(0), int64(32), "", nil},

		{SomeStructWithInt32WithErrFunc{}, int64(0), int64(32), "", nil},
		{&SomeStructWithInt32WithErrFunc{}, int64(0), int64(32), "", nil},

		{SomeStructWithInt32WithErrFuncPtr{}, int64(0), int64(32), "", nil},
		{&SomeStructWithInt32WithErrFuncPtr{}, int64(0), int64(32), "", nil},

		{SomeStructWithInt64Func{}, int64(0), int64(64), "", nil},
		{&SomeStructWithInt64Func{}, int64(0), int64(64), "", nil},

		{SomeStructWithInt64FuncPtr{}, int64(0), int64(64), "", nil},
		{&SomeStructWithInt64FuncPtr{}, int64(0), int64(64), "", nil},

		{SomeStructWithInt64WithErrFunc{}, int64(0), int64(64), "", nil},
		{&SomeStructWithInt64WithErrFunc{}, int64(0), int64(64), "", nil},

		{SomeStructWithInt64WithErrFuncPtr{}, int64(0), int64(64), "", nil},
		{&SomeStructWithInt64WithErrFuncPtr{}, int64(0), int64(64), "", nil},

		{SomeStructWithUintFunc{}, int64(0), int64(16), "", nil},
		{&SomeStructWithUintFunc{}, int64(0), int64(16), "", nil},

		{SomeStructWithUintFuncPtr{}, int64(0), int64(16), "", nil},
		{&SomeStructWithUintFuncPtr{}, int64(0), int64(16), "", nil},

		{SomeStructWithUintWithErrFunc{}, int64(0), int64(16), "", nil},
		{&SomeStructWithUintWithErrFunc{}, int64(0), int64(16), "", nil},

		{SomeStructWithUintWithErrFuncPtr{}, int64(0), int64(16), "", nil},
		{&SomeStructWithUintWithErrFuncPtr{}, int64(0), int64(16), "", nil},

		{SomeStructWithUint8Func{}, int64(0), int64(8), "", nil},
		{&SomeStructWithUint8Func{}, int64(0), int64(8), "", nil},

		{SomeStructWithUint8FuncPtr{}, int64(0), int64(8), "", nil},
		{&SomeStructWithUint8FuncPtr{}, int64(0), int64(8), "", nil},

		{SomeStructWithUint8WithErrFunc{}, int64(0), int64(8), "", nil},
		{&SomeStructWithUint8WithErrFunc{}, int64(0), int64(8), "", nil},

		{SomeStructWithUint8WithErrFuncPtr{}, int64(0), int64(8), "", nil},
		{&SomeStructWithUint8WithErrFuncPtr{}, int64(0), int64(8), "", nil},

		{SomeStructWithUint16Func{}, int64(0), int64(16), "", nil},
		{&SomeStructWithUint16Func{}, int64(0), int64(16), "", nil},

		{SomeStructWithUint16FuncPtr{}, int64(0), int64(16), "", nil},
		{&SomeStructWithUint16FuncPtr{}, int64(0), int64(16), "", nil},

		{SomeStructWithUint16WithErrFunc{}, int64(0), int64(16), "", nil},
		{&SomeStructWithUint16WithErrFunc{}, int64(0), int64(16), "", nil},

		{SomeStructWithUint16WithErrFuncPtr{}, int64(0), int64(16), "", nil},
		{&SomeStructWithUint16WithErrFuncPtr{}, int64(0), int64(16), "", nil},

		{SomeStructWithUint32Func{}, int64(0), int64(32), "", nil},
		{&SomeStructWithUint32Func{}, int64(0), int64(32), "", nil},

		{SomeStructWithUint32FuncPtr{}, int64(0), int64(32), "", nil},
		{&SomeStructWithUint32FuncPtr{}, int64(0), int64(32), "", nil},

		{SomeStructWithUint32WithErrFunc{}, int64(0), int64(32), "", nil},
		{&SomeStructWithUint32WithErrFunc{}, int64(0), int64(32), "", nil},

		{SomeStructWithUint32WithErrFuncPtr{}, int64(0), int64(32), "", nil},
		{&SomeStructWithUint32WithErrFuncPtr{}, int64(0), int64(32), "", nil},

		{SomeStructWithUint64Func{}, int64(0), int64(64), "", nil},
		{&SomeStructWithUint64Func{}, int64(0), int64(64), "", nil},

		{SomeStructWithUint64FuncPtr{}, int64(0), int64(64), "", nil},
		{&SomeStructWithUint64FuncPtr{}, int64(0), int64(64), "", nil},

		{SomeStructWithUint64WithErrFunc{}, int64(0), int64(64), "", nil},
		{&SomeStructWithUint64WithErrFunc{}, int64(0), int64(64), "", nil},

		{SomeStructWithUint64WithErrFuncPtr{}, int64(0), int64(64), "", nil},
		{&SomeStructWithUint64WithErrFuncPtr{}, int64(0), int64(64), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
