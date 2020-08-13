package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithFloat32Func struct{}

func (SomeStructWithFloat32Func) Float32() float32 {
	return 10
}

type SomeStructWithFloat32FuncPtr struct{}

func (*SomeStructWithFloat32FuncPtr) Float32() float32 {
	return 10
}

type SomeStructWithFloat32WithErrFunc struct{}

func (SomeStructWithFloat32WithErrFunc) Float32() (float32, error) {
	return 10, nil
}

type SomeStructWithFloat32WithErrFuncPtr struct{}

func (*SomeStructWithFloat32WithErrFuncPtr) Float32() (float32, error) {
	return 10, nil
}

func TestFloat32(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, float32(0), float32(0), "", nil},
		// string
		{"3.2", float32(0), float32(3.2), "", nil},
		{"Hello World", float32(0), float32(0), `unable to convert string to float32: strconv.ParseFloat: parsing "Hello World": invalid syntax`, nil},
		{"", float32(0), float32(0), "", nil},
		// bool
		{true, float32(0), float32(1), "", nil},
		{false, float32(0), float32(0), "", nil},
		// int
		{6, float32(0), float32(6), "", nil},
		// int8
		{int8(6), float32(0), float32(6), "", nil},
		// int16
		{int16(6), float32(0), float32(6), "", nil},
		// int32
		{int32(6), float32(0), float32(6), "", nil},
		// int64
		{int64(6), float32(0), float32(6), "", nil},
		// uint
		{uint(6), float32(0), float32(6), "", nil},
		// uint8
		{uint8(6), float32(0), float32(6), "", nil},
		// uint16
		{uint16(6), float32(0), float32(6), "", nil},
		// uint32
		{uint32(6), float32(0), float32(6), "", nil},
		// uint64
		{uint64(6), float32(0), float32(6), "", nil},
		// float32
		{float32(6), float32(0), float32(6), "", nil},
		// float64
		{float64(6), float32(0), float32(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, float32(0), float32(0), "unable to convert []int to float32: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, float32(0), float32(0), "unable to convert []uint8 to float32: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, float32(0), float32(0), "unable to convert []int32 to float32: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, float32(0), float32(0), "unable to convert []string to float32: no recipe", nil},
		// struct
		{struct{}{}, float32(0), float32(0), "unable to convert struct {} to float32: struct {} has no Float32() function", nil},
		// time
		{time.Unix(10, 10), float32(10), float32(10), "", nil},

		{SomeStructWithFloat32Func{}, float32(0), float32(10), "", nil},
		{&SomeStructWithFloat32Func{}, float32(0), float32(10), "", nil},

		{SomeStructWithFloat32FuncPtr{}, float32(0), float32(10), "", nil},
		{&SomeStructWithFloat32FuncPtr{}, float32(0), float32(10), "", nil},

		{SomeStructWithFloat32WithErrFunc{}, float32(0), float32(10), "", nil},
		{&SomeStructWithFloat32WithErrFunc{}, float32(0), float32(10), "", nil},

		{SomeStructWithFloat32WithErrFuncPtr{}, float32(0), float32(10), "", nil},
		{&SomeStructWithFloat32WithErrFuncPtr{}, float32(0), float32(10), "", nil},

		{SomeStructWithFloat64Func{}, float32(0), float32(10), "", nil},
		{&SomeStructWithFloat64Func{}, float32(0), float32(10), "", nil},

		{SomeStructWithFloat64FuncPtr{}, float32(0), float32(10), "", nil},
		{&SomeStructWithFloat64FuncPtr{}, float32(0), float32(10), "", nil},

		{SomeStructWithFloat64WithErrFunc{}, float32(0), float32(10), "", nil},
		{&SomeStructWithFloat64WithErrFunc{}, float32(0), float32(10), "", nil},

		{SomeStructWithFloat64WithErrFuncPtr{}, float32(0), float32(10), "", nil},
		{&SomeStructWithFloat64WithErrFuncPtr{}, float32(0), float32(10), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
