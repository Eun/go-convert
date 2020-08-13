package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithUint16Func struct{}

func (SomeStructWithUint16Func) Uint16() uint16 {
	return 16
}

type SomeStructWithUint16FuncPtr struct{}

func (*SomeStructWithUint16FuncPtr) Uint16() uint16 {
	return 16
}

type SomeStructWithUint16WithErrFunc struct{}

func (SomeStructWithUint16WithErrFunc) Uint16() (uint16, error) {
	return 16, nil
}

type SomeStructWithUint16WithErrFuncPtr struct{}

func (*SomeStructWithUint16WithErrFuncPtr) Uint16() (uint16, error) {
	return 16, nil
}

func TestUint16(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, uint16(0), uint16(0), "", nil},
		// string
		{"6", uint16(0), uint16(6), "", nil},
		{"", uint16(0), uint16(0), "", nil},
		{"Hello World", uint16(0), uint16(0), `unable to convert string to uint16: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint16(0), uint16(1), "", nil},
		{false, uint16(0), uint16(0), "", nil},
		// int
		{6, uint16(0), uint16(6), "", nil},
		// int8
		{int8(6), uint16(0), uint16(6), "", nil},
		// int16
		{int16(6), uint16(0), uint16(6), "", nil},
		// int32
		{int32(6), uint16(0), uint16(6), "", nil},
		// int64
		{int64(6), uint16(0), uint16(6), "", nil},
		// uint
		{uint(6), uint16(0), uint16(6), "", nil},
		// uint8
		{uint8(6), uint16(0), uint16(6), "", nil},
		// uint16
		{uint16(6), uint16(0), uint16(6), "", nil},
		// uint32
		{uint32(6), uint16(0), uint16(6), "", nil},
		// uint64
		{uint64(6), uint16(0), uint16(6), "", nil},
		// float32
		{float32(6), uint16(0), uint16(6), "", nil},
		// float64
		{float64(6), uint16(0), uint16(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint16(0), uint16(0), "unable to convert []int to uint16: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint16(0), uint16(0), "unable to convert []uint8 to uint16: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint16(0), uint16(0), "unable to convert []int32 to uint16: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint16(0), uint16(0), "unable to convert []string to uint16: no recipe", nil},
		// struct
		{struct{}{}, uint16(0), uint16(0), "unable to convert struct {} to uint16: struct {} has no Uint16() function", nil},
		// time
		{time.Unix(10, 10), uint16(10), uint16(10), "", nil},

		{SomeStructWithIntFunc{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithIntFunc{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithIntFuncPtr{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithIntFuncPtr{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithIntWithErrFunc{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithIntWithErrFunc{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithIntWithErrFuncPtr{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithIntWithErrFuncPtr{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithInt8Func{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithInt8Func{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithInt8FuncPtr{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithInt8FuncPtr{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithInt8WithErrFunc{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithInt8WithErrFunc{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithInt8WithErrFuncPtr{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithInt8WithErrFuncPtr{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithInt16Func{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithInt16Func{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithInt16FuncPtr{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithInt16FuncPtr{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithInt16WithErrFunc{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithInt16WithErrFunc{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithInt16WithErrFuncPtr{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithInt16WithErrFuncPtr{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithInt32Func{}, uint16(0), uint16(32), "", nil},
		{&SomeStructWithInt32Func{}, uint16(0), uint16(32), "", nil},

		{SomeStructWithInt32FuncPtr{}, uint16(0), uint16(32), "", nil},
		{&SomeStructWithInt32FuncPtr{}, uint16(0), uint16(32), "", nil},

		{SomeStructWithInt32WithErrFunc{}, uint16(0), uint16(32), "", nil},
		{&SomeStructWithInt32WithErrFunc{}, uint16(0), uint16(32), "", nil},

		{SomeStructWithInt32WithErrFuncPtr{}, uint16(0), uint16(32), "", nil},
		{&SomeStructWithInt32WithErrFuncPtr{}, uint16(0), uint16(32), "", nil},

		{SomeStructWithInt64Func{}, uint16(0), uint16(64), "", nil},
		{&SomeStructWithInt64Func{}, uint16(0), uint16(64), "", nil},

		{SomeStructWithInt64FuncPtr{}, uint16(0), uint16(64), "", nil},
		{&SomeStructWithInt64FuncPtr{}, uint16(0), uint16(64), "", nil},

		{SomeStructWithInt64WithErrFunc{}, uint16(0), uint16(64), "", nil},
		{&SomeStructWithInt64WithErrFunc{}, uint16(0), uint16(64), "", nil},

		{SomeStructWithInt64WithErrFuncPtr{}, uint16(0), uint16(64), "", nil},
		{&SomeStructWithInt64WithErrFuncPtr{}, uint16(0), uint16(64), "", nil},

		{SomeStructWithUintFunc{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithUintFunc{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithUintFuncPtr{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithUintFuncPtr{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithUintWithErrFunc{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithUintWithErrFunc{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithUintWithErrFuncPtr{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithUintWithErrFuncPtr{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithUint8Func{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithUint8Func{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithUint8FuncPtr{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithUint8FuncPtr{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithUint8WithErrFunc{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithUint8WithErrFunc{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithUint8WithErrFuncPtr{}, uint16(0), uint16(8), "", nil},
		{&SomeStructWithUint8WithErrFuncPtr{}, uint16(0), uint16(8), "", nil},

		{SomeStructWithUint16Func{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithUint16Func{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithUint16FuncPtr{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithUint16FuncPtr{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithUint16WithErrFunc{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithUint16WithErrFunc{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithUint16WithErrFuncPtr{}, uint16(0), uint16(16), "", nil},
		{&SomeStructWithUint16WithErrFuncPtr{}, uint16(0), uint16(16), "", nil},

		{SomeStructWithUint32Func{}, uint16(0), uint16(32), "", nil},
		{&SomeStructWithUint32Func{}, uint16(0), uint16(32), "", nil},

		{SomeStructWithUint32FuncPtr{}, uint16(0), uint16(32), "", nil},
		{&SomeStructWithUint32FuncPtr{}, uint16(0), uint16(32), "", nil},

		{SomeStructWithUint32WithErrFunc{}, uint16(0), uint16(32), "", nil},
		{&SomeStructWithUint32WithErrFunc{}, uint16(0), uint16(32), "", nil},

		{SomeStructWithUint32WithErrFuncPtr{}, uint16(0), uint16(32), "", nil},
		{&SomeStructWithUint32WithErrFuncPtr{}, uint16(0), uint16(32), "", nil},

		{SomeStructWithUint64Func{}, uint16(0), uint16(64), "", nil},
		{&SomeStructWithUint64Func{}, uint16(0), uint16(64), "", nil},

		{SomeStructWithUint64FuncPtr{}, uint16(0), uint16(64), "", nil},
		{&SomeStructWithUint64FuncPtr{}, uint16(0), uint16(64), "", nil},

		{SomeStructWithUint64WithErrFunc{}, uint16(0), uint16(64), "", nil},
		{&SomeStructWithUint64WithErrFunc{}, uint16(0), uint16(64), "", nil},

		{SomeStructWithUint64WithErrFuncPtr{}, uint16(0), uint16(64), "", nil},
		{&SomeStructWithUint64WithErrFuncPtr{}, uint16(0), uint16(64), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
