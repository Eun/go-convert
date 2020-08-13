package convert_test

import (
	"testing"
	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithIntFunc struct{}

func (SomeStructWithIntFunc) Int() int {
	return 8
}

type SomeStructWithIntFuncPtr struct{}

func (*SomeStructWithIntFuncPtr) Int() int {
	return 8
}

type SomeStructWithIntWithErrFunc struct{}

func (SomeStructWithIntWithErrFunc) Int() (int, error) {
	return 8, nil
}

type SomeStructWithIntWithErrFuncPtr struct{}

func (*SomeStructWithIntWithErrFuncPtr) Int() (int, error) {
	return 8, nil
}

func TestInt(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, 0, 0, "", nil},
		// string
		{"6", 0, 6, "", nil},
		{"", 0, 0, "", nil},
		{"Hello World", 0, 0, `unable to convert string to int: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, 0, 1, "", nil},
		{false, 0, 0, "", nil},
		// int
		{6, 0, 6, "", nil},
		// int8
		{int8(6), 0, 6, "", nil},
		// int16
		{int16(6), 0, 6, "", nil},
		// int32
		{int32(6), 0, 6, "", nil},
		// int64
		{int64(6), 0, 6, "", nil},
		// uint
		{uint(6), 0, 6, "", nil},
		// uint8
		{uint8(6), 0, 6, "", nil},
		// uint16
		{uint16(6), 0, 6, "", nil},
		// uint32
		{uint32(6), 0, 6, "", nil},
		// uint64
		{uint64(6), 0, 6, "", nil},
		// float32
		{float32(6), 0, 6, "", nil},
		// float64
		{float64(6), 0, 6, "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, 0, 0, "unable to convert []int to int: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, 0, 0, "unable to convert []uint8 to int: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, 0, 0, "unable to convert []int32 to int: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, 0, 0, "unable to convert []string to int: no recipe", nil},
		// struct
		{struct{}{}, 0, 0, "unable to convert struct {} to int: struct {} has no Int() function", nil},
		// time
		{time.Unix(10, 10), int(10), int(10), "", nil},

		{SomeStructWithIntFunc{}, int(0), int(8), "", nil},
		{&SomeStructWithIntFunc{}, int(0), int(8), "", nil},

		{SomeStructWithIntFuncPtr{}, int(0), int(8), "", nil},
		{&SomeStructWithIntFuncPtr{}, int(0), int(8), "", nil},

		{SomeStructWithIntWithErrFunc{}, int(0), int(8), "", nil},
		{&SomeStructWithIntWithErrFunc{}, int(0), int(8), "", nil},

		{SomeStructWithIntWithErrFuncPtr{}, int(0), int(8), "", nil},
		{&SomeStructWithIntWithErrFuncPtr{}, int(0), int(8), "", nil},

		{SomeStructWithInt8Func{}, int(0), int(8), "", nil},
		{&SomeStructWithInt8Func{}, int(0), int(8), "", nil},

		{SomeStructWithInt8FuncPtr{}, int(0), int(8), "", nil},
		{&SomeStructWithInt8FuncPtr{}, int(0), int(8), "", nil},

		{SomeStructWithInt8WithErrFunc{}, int(0), int(8), "", nil},
		{&SomeStructWithInt8WithErrFunc{}, int(0), int(8), "", nil},

		{SomeStructWithInt8WithErrFuncPtr{}, int(0), int(8), "", nil},
		{&SomeStructWithInt8WithErrFuncPtr{}, int(0), int(8), "", nil},

		{SomeStructWithInt16Func{}, int(0), int(16), "", nil},
		{&SomeStructWithInt16Func{}, int(0), int(16), "", nil},

		{SomeStructWithInt16FuncPtr{}, int(0), int(16), "", nil},
		{&SomeStructWithInt16FuncPtr{}, int(0), int(16), "", nil},

		{SomeStructWithInt16WithErrFunc{}, int(0), int(16), "", nil},
		{&SomeStructWithInt16WithErrFunc{}, int(0), int(16), "", nil},

		{SomeStructWithInt16WithErrFuncPtr{}, int(0), int(16), "", nil},
		{&SomeStructWithInt16WithErrFuncPtr{}, int(0), int(16), "", nil},

		{SomeStructWithInt32Func{}, int(0), int(32), "", nil},
		{&SomeStructWithInt32Func{}, int(0), int(32), "", nil},

		{SomeStructWithInt32FuncPtr{}, int(0), int(32), "", nil},
		{&SomeStructWithInt32FuncPtr{}, int(0), int(32), "", nil},

		{SomeStructWithInt32WithErrFunc{}, int(0), int(32), "", nil},
		{&SomeStructWithInt32WithErrFunc{}, int(0), int(32), "", nil},

		{SomeStructWithInt32WithErrFuncPtr{}, int(0), int(32), "", nil},
		{&SomeStructWithInt32WithErrFuncPtr{}, int(0), int(32), "", nil},

		{SomeStructWithInt64Func{}, int(0), int(64), "", nil},
		{&SomeStructWithInt64Func{}, int(0), int(64), "", nil},

		{SomeStructWithInt64FuncPtr{}, int(0), int(64), "", nil},
		{&SomeStructWithInt64FuncPtr{}, int(0), int(64), "", nil},

		{SomeStructWithInt64WithErrFunc{}, int(0), int(64), "", nil},
		{&SomeStructWithInt64WithErrFunc{}, int(0), int(64), "", nil},

		{SomeStructWithInt64WithErrFuncPtr{}, int(0), int(64), "", nil},
		{&SomeStructWithInt64WithErrFuncPtr{}, int(0), int(64), "", nil},

		{SomeStructWithUintFunc{}, int(0), int(16), "", nil},
		{&SomeStructWithUintFunc{}, int(0), int(16), "", nil},

		{SomeStructWithUintFuncPtr{}, int(0), int(16), "", nil},
		{&SomeStructWithUintFuncPtr{}, int(0), int(16), "", nil},

		{SomeStructWithUintWithErrFunc{}, int(0), int(16), "", nil},
		{&SomeStructWithUintWithErrFunc{}, int(0), int(16), "", nil},

		{SomeStructWithUintWithErrFuncPtr{}, int(0), int(16), "", nil},
		{&SomeStructWithUintWithErrFuncPtr{}, int(0), int(16), "", nil},

		{SomeStructWithUint8Func{}, int(0), int(8), "", nil},
		{&SomeStructWithUint8Func{}, int(0), int(8), "", nil},

		{SomeStructWithUint8FuncPtr{}, int(0), int(8), "", nil},
		{&SomeStructWithUint8FuncPtr{}, int(0), int(8), "", nil},

		{SomeStructWithUint8WithErrFunc{}, int(0), int(8), "", nil},
		{&SomeStructWithUint8WithErrFunc{}, int(0), int(8), "", nil},

		{SomeStructWithUint8WithErrFuncPtr{}, int(0), int(8), "", nil},
		{&SomeStructWithUint8WithErrFuncPtr{}, int(0), int(8), "", nil},

		{SomeStructWithUint16Func{}, int(0), int(16), "", nil},
		{&SomeStructWithUint16Func{}, int(0), int(16), "", nil},

		{SomeStructWithUint16FuncPtr{}, int(0), int(16), "", nil},
		{&SomeStructWithUint16FuncPtr{}, int(0), int(16), "", nil},

		{SomeStructWithUint16WithErrFunc{}, int(0), int(16), "", nil},
		{&SomeStructWithUint16WithErrFunc{}, int(0), int(16), "", nil},

		{SomeStructWithUint16WithErrFuncPtr{}, int(0), int(16), "", nil},
		{&SomeStructWithUint16WithErrFuncPtr{}, int(0), int(16), "", nil},

		{SomeStructWithUint32Func{}, int(0), int(32), "", nil},
		{&SomeStructWithUint32Func{}, int(0), int(32), "", nil},

		{SomeStructWithUint32FuncPtr{}, int(0), int(32), "", nil},
		{&SomeStructWithUint32FuncPtr{}, int(0), int(32), "", nil},

		{SomeStructWithUint32WithErrFunc{}, int(0), int(32), "", nil},
		{&SomeStructWithUint32WithErrFunc{}, int(0), int(32), "", nil},

		{SomeStructWithUint32WithErrFuncPtr{}, int(0), int(32), "", nil},
		{&SomeStructWithUint32WithErrFuncPtr{}, int(0), int(32), "", nil},

		{SomeStructWithUint64Func{}, int(0), int(64), "", nil},
		{&SomeStructWithUint64Func{}, int(0), int(64), "", nil},

		{SomeStructWithUint64FuncPtr{}, int(0), int(64), "", nil},
		{&SomeStructWithUint64FuncPtr{}, int(0), int(64), "", nil},

		{SomeStructWithUint64WithErrFunc{}, int(0), int(64), "", nil},
		{&SomeStructWithUint64WithErrFunc{}, int(0), int(64), "", nil},

		{SomeStructWithUint64WithErrFuncPtr{}, int(0), int(64), "", nil},
		{&SomeStructWithUint64WithErrFuncPtr{}, int(0), int(64), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
