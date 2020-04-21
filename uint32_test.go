package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithUint32Func struct{}

func (SomeStructWithUint32Func) Uint32() uint32 {
	return 32
}

type SomeStructWithUint32FuncPtr struct{}

func (*SomeStructWithUint32FuncPtr) Uint32() uint32 {
	return 32
}

type SomeStructWithUint32WithErrFunc struct{}

func (SomeStructWithUint32WithErrFunc) Uint32() (uint32, error) {
	return 32, nil
}

type SomeStructWithUint32WithErrFuncPtr struct{}

func (*SomeStructWithUint32WithErrFuncPtr) Uint32() (uint32, error) {
	return 32, nil
}

func TestUint32(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, uint32(0), uint32(0), "", nil},
		// string
		{"6", uint32(0), uint32(6), "", nil},
		{"", uint32(0), uint32(0), "", nil},
		{"Hello World", uint32(0), uint32(0), `unable to convert string to uint32: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint32(0), uint32(1), "", nil},
		{false, uint32(0), uint32(0), "", nil},
		// int
		{6, uint32(0), uint32(6), "", nil},
		// int8
		{int8(6), uint32(0), uint32(6), "", nil},
		// int16
		{int16(6), uint32(0), uint32(6), "", nil},
		// int32
		{int32(6), uint32(0), uint32(6), "", nil},
		// int64
		{int64(6), uint32(0), uint32(6), "", nil},
		// uint
		{uint(6), uint32(0), uint32(6), "", nil},
		// uint8
		{uint8(6), uint32(0), uint32(6), "", nil},
		// uint16
		{uint16(6), uint32(0), uint32(6), "", nil},
		// uint32
		{uint32(6), uint32(0), uint32(6), "", nil},
		// uint64
		{uint64(6), uint32(0), uint32(6), "", nil},
		// float32
		{float32(6), uint32(0), uint32(6), "", nil},
		// float64
		{float64(6), uint32(0), uint32(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint32(0), uint32(0), "unable to convert []int to uint32: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint32(0), uint32(0), "unable to convert []uint8 to uint32: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint32(0), uint32(0), "unable to convert []int32 to uint32: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint32(0), uint32(0), "unable to convert []string to uint32: no recipe", nil},
		// struct
		{struct{}{}, uint32(0), uint32(0), "unable to convert struct {} to uint32: struct {} has no Uint32() function", nil},
		// time
		{time.Unix(10, 10), uint32(10), uint32(10), "", nil},

		{SomeStructWithUint32Func{}, uint32(0), uint32(32), "", nil},
		{&SomeStructWithUint32Func{}, uint32(0), uint32(32), "", nil},

		{SomeStructWithUint32FuncPtr{}, uint32(0), uint32(32), "", nil},
		{&SomeStructWithUint32FuncPtr{}, uint32(0), uint32(32), "", nil},

		{SomeStructWithUint32WithErrFunc{}, uint32(0), uint32(32), "", nil},
		{&SomeStructWithUint32WithErrFunc{}, uint32(0), uint32(32), "", nil},

		{SomeStructWithUint32WithErrFuncPtr{}, uint32(0), uint32(32), "", nil},
		{&SomeStructWithUint32WithErrFuncPtr{}, uint32(0), uint32(32), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
