package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithUint64Func struct{}

func (SomeStructWithUint64Func) Uint64() uint64 {
	return 64
}

type SomeStructWithUint64FuncPtr struct{}

func (*SomeStructWithUint64FuncPtr) Uint64() uint64 {
	return 64
}

type SomeStructWithUint64WithErrFunc struct{}

func (SomeStructWithUint64WithErrFunc) Uint64() (uint64, error) {
	return 64, nil
}

type SomeStructWithUint64WithErrFuncPtr struct{}

func (*SomeStructWithUint64WithErrFuncPtr) Uint64() (uint64, error) {
	return 64, nil
}

func TestUint64(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, uint64(0), uint64(0), "", nil},
		// string
		{"6", uint64(0), uint64(6), "", nil},
		{"", uint64(0), uint64(0), "", nil},
		{"Hello World", uint64(0), uint64(0), `unable to convert string to uint64: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint64(0), uint64(1), "", nil},
		{false, uint64(0), uint64(0), "", nil},
		// int
		{6, uint64(0), uint64(6), "", nil},
		// int8
		{int8(6), uint64(0), uint64(6), "", nil},
		// int16
		{int16(6), uint64(0), uint64(6), "", nil},
		// int32
		{int32(6), uint64(0), uint64(6), "", nil},
		// int64
		{int64(6), uint64(0), uint64(6), "", nil},
		// uint
		{uint(6), uint64(0), uint64(6), "", nil},
		// uint8
		{uint8(6), uint64(0), uint64(6), "", nil},
		// uint16
		{uint16(6), uint64(0), uint64(6), "", nil},
		// uint32
		{uint32(6), uint64(0), uint64(6), "", nil},
		// uint64
		{uint64(6), uint64(0), uint64(6), "", nil},
		// float32
		{float32(6), uint64(0), uint64(6), "", nil},
		// float64
		{float64(6), uint64(0), uint64(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint64(0), uint64(0), "unable to convert []int to uint64: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint64(0), uint64(0), "unable to convert []uint8 to uint64: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint64(0), uint64(0), "unable to convert []int32 to uint64: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint64(0), uint64(0), "unable to convert []string to uint64: no recipe", nil},
		// struct
		{struct{}{}, uint64(0), uint64(0), "unable to convert struct {} to uint64: struct {} has no Uint64() function", nil},
		// time
		{time.Unix(10, 10), uint64(10), uint64(10), "", nil},

		{SomeStructWithUint64Func{}, uint64(0), uint64(64), "", nil},
		{&SomeStructWithUint64FuncPtr{}, uint64(0), uint64(64), "", nil},
		{SomeStructWithUint64WithErrFunc{}, uint64(0), uint64(64), "", nil},
		{&SomeStructWithUint64WithErrFuncPtr{}, uint64(0), uint64(64), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
