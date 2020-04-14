package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithInt8Func struct {
}

func (SomeStructWithInt8Func) Int8() int8 {
	return 8
}

type SomeStructWithInt8FuncPtr struct {
}

func (*SomeStructWithInt8FuncPtr) Int8() int8 {
	return 8
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

		{SomeStructWithInt8Func{}, int8(0), int8(8), "", nil},
		{&SomeStructWithInt8FuncPtr{}, int8(0), int8(8), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
