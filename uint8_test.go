package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithUint8Func struct{}

func (SomeStructWithUint8Func) Uint8() uint8 {
	return 8
}

type SomeStructWithUint8FuncPtr struct{}

func (*SomeStructWithUint8FuncPtr) Uint8() uint8 {
	return 8
}

type SomeStructWithUint8WithErrFunc struct{}

func (SomeStructWithUint8WithErrFunc) Uint8() (uint8, error) {
	return 8, nil
}

type SomeStructWithUint8WithErrFuncPtr struct{}

func (*SomeStructWithUint8WithErrFuncPtr) Uint8() (uint8, error) {
	return 8, nil
}

func TestUint8(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, uint8(0), uint8(0), "", nil},
		// string
		{"6", uint8(0), uint8(6), "", nil},
		{"", uint8(0), uint8(0), "", nil},
		{"Hello World", uint8(0), uint8(0), `unable to convert string to uint8: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint8(0), uint8(1), "", nil},
		{false, uint8(0), uint8(0), "", nil},
		// int
		{6, uint8(0), uint8(6), "", nil},
		// int8
		{int8(6), uint8(0), uint8(6), "", nil},
		// int16
		{int16(6), uint8(0), uint8(6), "", nil},
		// int32
		{int32(6), uint8(0), uint8(6), "", nil},
		// int64
		{int64(6), uint8(0), uint8(6), "", nil},
		// uint
		{uint(6), uint8(0), uint8(6), "", nil},
		// int8
		{uint8(6), uint8(0), uint8(6), "", nil},
		// uint16
		{uint16(6), uint8(0), uint8(6), "", nil},
		// uint32
		{uint32(6), uint8(0), uint8(6), "", nil},
		// uint64
		{uint64(6), uint8(0), uint8(6), "", nil},
		// float32
		{float32(6), uint8(0), uint8(6), "", nil},
		// float64
		{float64(6), uint8(0), uint8(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint8(0), uint8(0), "unable to convert []int to uint8: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint8(0), uint8(0), "unable to convert []uint8 to uint8: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint8(0), uint8(0), "unable to convert []int32 to uint8: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint8(0), uint8(0), "unable to convert []string to uint8: no recipe", nil},
		// struct
		{struct{}{}, uint8(0), uint8(0), "unable to convert struct {} to uint8: struct {} has no Uint8() function", nil},
		// time
		{time.Unix(10, 10), uint8(10), uint8(10), "", nil},

		{SomeStructWithIntFunc{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithIntFunc{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithIntFuncPtr{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithIntFuncPtr{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithIntWithErrFunc{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithIntWithErrFunc{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithIntWithErrFuncPtr{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithIntWithErrFuncPtr{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithInt8Func{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithInt8Func{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithInt8FuncPtr{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithInt8FuncPtr{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithInt8WithErrFunc{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithInt8WithErrFunc{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithInt8WithErrFuncPtr{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithInt8WithErrFuncPtr{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithInt16Func{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithInt16Func{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithInt16FuncPtr{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithInt16FuncPtr{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithInt16WithErrFunc{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithInt16WithErrFunc{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithInt16WithErrFuncPtr{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithInt16WithErrFuncPtr{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithInt32Func{}, uint8(0), uint8(32), "", nil},
		{&SomeStructWithInt32Func{}, uint8(0), uint8(32), "", nil},

		{SomeStructWithInt32FuncPtr{}, uint8(0), uint8(32), "", nil},
		{&SomeStructWithInt32FuncPtr{}, uint8(0), uint8(32), "", nil},

		{SomeStructWithInt32WithErrFunc{}, uint8(0), uint8(32), "", nil},
		{&SomeStructWithInt32WithErrFunc{}, uint8(0), uint8(32), "", nil},

		{SomeStructWithInt32WithErrFuncPtr{}, uint8(0), uint8(32), "", nil},
		{&SomeStructWithInt32WithErrFuncPtr{}, uint8(0), uint8(32), "", nil},

		{SomeStructWithInt64Func{}, uint8(0), uint8(64), "", nil},
		{&SomeStructWithInt64Func{}, uint8(0), uint8(64), "", nil},

		{SomeStructWithInt64FuncPtr{}, uint8(0), uint8(64), "", nil},
		{&SomeStructWithInt64FuncPtr{}, uint8(0), uint8(64), "", nil},

		{SomeStructWithInt64WithErrFunc{}, uint8(0), uint8(64), "", nil},
		{&SomeStructWithInt64WithErrFunc{}, uint8(0), uint8(64), "", nil},

		{SomeStructWithInt64WithErrFuncPtr{}, uint8(0), uint8(64), "", nil},
		{&SomeStructWithInt64WithErrFuncPtr{}, uint8(0), uint8(64), "", nil},

		{SomeStructWithUintFunc{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithUintFunc{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithUintFuncPtr{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithUintFuncPtr{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithUintWithErrFunc{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithUintWithErrFunc{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithUintWithErrFuncPtr{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithUintWithErrFuncPtr{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithUint8Func{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithUint8Func{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithUint8FuncPtr{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithUint8FuncPtr{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithUint8WithErrFunc{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithUint8WithErrFunc{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithUint8WithErrFuncPtr{}, uint8(0), uint8(8), "", nil},
		{&SomeStructWithUint8WithErrFuncPtr{}, uint8(0), uint8(8), "", nil},

		{SomeStructWithUint16Func{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithUint16Func{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithUint16FuncPtr{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithUint16FuncPtr{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithUint16WithErrFunc{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithUint16WithErrFunc{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithUint16WithErrFuncPtr{}, uint8(0), uint8(16), "", nil},
		{&SomeStructWithUint16WithErrFuncPtr{}, uint8(0), uint8(16), "", nil},

		{SomeStructWithUint32Func{}, uint8(0), uint8(32), "", nil},
		{&SomeStructWithUint32Func{}, uint8(0), uint8(32), "", nil},

		{SomeStructWithUint32FuncPtr{}, uint8(0), uint8(32), "", nil},
		{&SomeStructWithUint32FuncPtr{}, uint8(0), uint8(32), "", nil},

		{SomeStructWithUint32WithErrFunc{}, uint8(0), uint8(32), "", nil},
		{&SomeStructWithUint32WithErrFunc{}, uint8(0), uint8(32), "", nil},

		{SomeStructWithUint32WithErrFuncPtr{}, uint8(0), uint8(32), "", nil},
		{&SomeStructWithUint32WithErrFuncPtr{}, uint8(0), uint8(32), "", nil},

		{SomeStructWithUint64Func{}, uint8(0), uint8(64), "", nil},
		{&SomeStructWithUint64Func{}, uint8(0), uint8(64), "", nil},

		{SomeStructWithUint64FuncPtr{}, uint8(0), uint8(64), "", nil},
		{&SomeStructWithUint64FuncPtr{}, uint8(0), uint8(64), "", nil},

		{SomeStructWithUint64WithErrFunc{}, uint8(0), uint8(64), "", nil},
		{&SomeStructWithUint64WithErrFunc{}, uint8(0), uint8(64), "", nil},

		{SomeStructWithUint64WithErrFuncPtr{}, uint8(0), uint8(64), "", nil},
		{&SomeStructWithUint64WithErrFuncPtr{}, uint8(0), uint8(64), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
