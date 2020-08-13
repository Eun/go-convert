package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithUintFunc struct{}

func (SomeStructWithUintFunc) Uint() uint {
	return 16
}

type SomeStructWithUintFuncPtr struct{}

func (*SomeStructWithUintFuncPtr) Uint() uint {
	return 16
}

type SomeStructWithUintWithErrFunc struct{}

func (SomeStructWithUintWithErrFunc) Uint() (uint, error) {
	return 16, nil
}

type SomeStructWithUintWithErrFuncPtr struct{}

func (*SomeStructWithUintWithErrFuncPtr) Uint() (uint, error) {
	return 16, nil
}

func TestUint(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, uint(0), uint(0), "", nil},
		// string
		{"6", uint(0), uint(6), "", nil},
		{"Hello World", uint(0), uint(0), `unable to convert string to uint: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint(0), uint(1), "", nil},
		{false, uint(0), uint(0), "", nil},
		// int
		{6, uint(0), uint(6), "", nil},
		// int8
		{int8(6), uint(0), uint(6), "", nil},
		// int16
		{int16(6), uint(0), uint(6), "", nil},
		// int32
		{int32(6), uint(0), uint(6), "", nil},
		// int64
		{int64(6), uint(0), uint(6), "", nil},
		// uint
		{uint(6), uint(0), uint(6), "", nil},
		// uint8
		{uint8(6), uint(0), uint(6), "", nil},
		// uint16
		{uint16(6), uint(0), uint(6), "", nil},
		// uint32
		{uint32(6), uint(0), uint(6), "", nil},
		// uint64
		{uint64(6), uint(0), uint(6), "", nil},
		// float32
		{float32(6), uint(0), uint(6), "", nil},
		// float64
		{float64(6), uint(0), uint(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint(0), uint(0), "unable to convert []int to uint: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint(0), uint(0), "unable to convert []uint8 to uint: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint(0), uint(0), "unable to convert []int32 to uint: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint(0), uint(0), "unable to convert []string to uint: no recipe", nil},
		// struct
		{struct{}{}, uint(0), uint(0), "unable to convert struct {} to uint: struct {} has no Uint() function", nil},
		// time
		{time.Unix(10, 10), uint(10), uint(10), "", nil},

		{SomeStructWithIntFunc{}, uint(0), uint(8), "", nil},
		{&SomeStructWithIntFunc{}, uint(0), uint(8), "", nil},

		{SomeStructWithIntFuncPtr{}, uint(0), uint(8), "", nil},
		{&SomeStructWithIntFuncPtr{}, uint(0), uint(8), "", nil},

		{SomeStructWithIntWithErrFunc{}, uint(0), uint(8), "", nil},
		{&SomeStructWithIntWithErrFunc{}, uint(0), uint(8), "", nil},

		{SomeStructWithIntWithErrFuncPtr{}, uint(0), uint(8), "", nil},
		{&SomeStructWithIntWithErrFuncPtr{}, uint(0), uint(8), "", nil},

		{SomeStructWithInt8Func{}, uint(0), uint(8), "", nil},
		{&SomeStructWithInt8Func{}, uint(0), uint(8), "", nil},

		{SomeStructWithInt8FuncPtr{}, uint(0), uint(8), "", nil},
		{&SomeStructWithInt8FuncPtr{}, uint(0), uint(8), "", nil},

		{SomeStructWithInt8WithErrFunc{}, uint(0), uint(8), "", nil},
		{&SomeStructWithInt8WithErrFunc{}, uint(0), uint(8), "", nil},

		{SomeStructWithInt8WithErrFuncPtr{}, uint(0), uint(8), "", nil},
		{&SomeStructWithInt8WithErrFuncPtr{}, uint(0), uint(8), "", nil},

		{SomeStructWithInt16Func{}, uint(0), uint(16), "", nil},
		{&SomeStructWithInt16Func{}, uint(0), uint(16), "", nil},

		{SomeStructWithInt16FuncPtr{}, uint(0), uint(16), "", nil},
		{&SomeStructWithInt16FuncPtr{}, uint(0), uint(16), "", nil},

		{SomeStructWithInt16WithErrFunc{}, uint(0), uint(16), "", nil},
		{&SomeStructWithInt16WithErrFunc{}, uint(0), uint(16), "", nil},

		{SomeStructWithInt16WithErrFuncPtr{}, uint(0), uint(16), "", nil},
		{&SomeStructWithInt16WithErrFuncPtr{}, uint(0), uint(16), "", nil},

		{SomeStructWithInt32Func{}, uint(0), uint(32), "", nil},
		{&SomeStructWithInt32Func{}, uint(0), uint(32), "", nil},

		{SomeStructWithInt32FuncPtr{}, uint(0), uint(32), "", nil},
		{&SomeStructWithInt32FuncPtr{}, uint(0), uint(32), "", nil},

		{SomeStructWithInt32WithErrFunc{}, uint(0), uint(32), "", nil},
		{&SomeStructWithInt32WithErrFunc{}, uint(0), uint(32), "", nil},

		{SomeStructWithInt32WithErrFuncPtr{}, uint(0), uint(32), "", nil},
		{&SomeStructWithInt32WithErrFuncPtr{}, uint(0), uint(32), "", nil},

		{SomeStructWithInt64Func{}, uint(0), uint(64), "", nil},
		{&SomeStructWithInt64Func{}, uint(0), uint(64), "", nil},

		{SomeStructWithInt64FuncPtr{}, uint(0), uint(64), "", nil},
		{&SomeStructWithInt64FuncPtr{}, uint(0), uint(64), "", nil},

		{SomeStructWithInt64WithErrFunc{}, uint(0), uint(64), "", nil},
		{&SomeStructWithInt64WithErrFunc{}, uint(0), uint(64), "", nil},

		{SomeStructWithInt64WithErrFuncPtr{}, uint(0), uint(64), "", nil},
		{&SomeStructWithInt64WithErrFuncPtr{}, uint(0), uint(64), "", nil},

		{SomeStructWithUintFunc{}, uint(0), uint(16), "", nil},
		{&SomeStructWithUintFunc{}, uint(0), uint(16), "", nil},

		{SomeStructWithUintFuncPtr{}, uint(0), uint(16), "", nil},
		{&SomeStructWithUintFuncPtr{}, uint(0), uint(16), "", nil},

		{SomeStructWithUintWithErrFunc{}, uint(0), uint(16), "", nil},
		{&SomeStructWithUintWithErrFunc{}, uint(0), uint(16), "", nil},

		{SomeStructWithUintWithErrFuncPtr{}, uint(0), uint(16), "", nil},
		{&SomeStructWithUintWithErrFuncPtr{}, uint(0), uint(16), "", nil},

		{SomeStructWithUint8Func{}, uint(0), uint(8), "", nil},
		{&SomeStructWithUint8Func{}, uint(0), uint(8), "", nil},

		{SomeStructWithUint8FuncPtr{}, uint(0), uint(8), "", nil},
		{&SomeStructWithUint8FuncPtr{}, uint(0), uint(8), "", nil},

		{SomeStructWithUint8WithErrFunc{}, uint(0), uint(8), "", nil},
		{&SomeStructWithUint8WithErrFunc{}, uint(0), uint(8), "", nil},

		{SomeStructWithUint8WithErrFuncPtr{}, uint(0), uint(8), "", nil},
		{&SomeStructWithUint8WithErrFuncPtr{}, uint(0), uint(8), "", nil},

		{SomeStructWithUint16Func{}, uint(0), uint(16), "", nil},
		{&SomeStructWithUint16Func{}, uint(0), uint(16), "", nil},

		{SomeStructWithUint16FuncPtr{}, uint(0), uint(16), "", nil},
		{&SomeStructWithUint16FuncPtr{}, uint(0), uint(16), "", nil},

		{SomeStructWithUint16WithErrFunc{}, uint(0), uint(16), "", nil},
		{&SomeStructWithUint16WithErrFunc{}, uint(0), uint(16), "", nil},

		{SomeStructWithUint16WithErrFuncPtr{}, uint(0), uint(16), "", nil},
		{&SomeStructWithUint16WithErrFuncPtr{}, uint(0), uint(16), "", nil},

		{SomeStructWithUint32Func{}, uint(0), uint(32), "", nil},
		{&SomeStructWithUint32Func{}, uint(0), uint(32), "", nil},

		{SomeStructWithUint32FuncPtr{}, uint(0), uint(32), "", nil},
		{&SomeStructWithUint32FuncPtr{}, uint(0), uint(32), "", nil},

		{SomeStructWithUint32WithErrFunc{}, uint(0), uint(32), "", nil},
		{&SomeStructWithUint32WithErrFunc{}, uint(0), uint(32), "", nil},

		{SomeStructWithUint32WithErrFuncPtr{}, uint(0), uint(32), "", nil},
		{&SomeStructWithUint32WithErrFuncPtr{}, uint(0), uint(32), "", nil},

		{SomeStructWithUint64Func{}, uint(0), uint(64), "", nil},
		{&SomeStructWithUint64Func{}, uint(0), uint(64), "", nil},

		{SomeStructWithUint64FuncPtr{}, uint(0), uint(64), "", nil},
		{&SomeStructWithUint64FuncPtr{}, uint(0), uint(64), "", nil},

		{SomeStructWithUint64WithErrFunc{}, uint(0), uint(64), "", nil},
		{&SomeStructWithUint64WithErrFunc{}, uint(0), uint(64), "", nil},

		{SomeStructWithUint64WithErrFuncPtr{}, uint(0), uint(64), "", nil},
		{&SomeStructWithUint64WithErrFuncPtr{}, uint(0), uint(64), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
