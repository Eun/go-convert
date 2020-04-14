package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithUintFunc struct {
}

func (SomeStructWithUintFunc) Uint() uint {
	return 16
}

type SomeStructWithUintFuncPtr struct {
}

func (*SomeStructWithUintFuncPtr) Uint() uint {
	return 16
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

		{SomeStructWithUintFunc{}, uint(0), uint(16), "", nil},
		{&SomeStructWithUintFuncPtr{}, uint(0), uint(16), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
