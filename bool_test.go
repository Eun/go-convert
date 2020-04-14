package convert_test

import (
	"testing"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithBoolFunc struct {
}

func (SomeStructWithBoolFunc) Bool() bool {
	return true
}

type SomeStructWithBoolFuncPtr struct {
}

func (*SomeStructWithBoolFuncPtr) Bool() bool {
	return true
}

func TestBool(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, false, false, "", nil},
		// string
		{"true", false, true, "", nil},
		{"Hello World", false, false, `unable to convert string to bool: strconv.ParseBool: parsing "Hello World": invalid syntax`, nil},
		{"", false, false, "", nil},
		// bool
		{true, false, true, "", nil},
		{false, false, false, "", nil},
		// int
		{6, false, true, "", nil},
		// int8
		{int8(6), false, true, "", nil},
		// int16
		{int16(6), false, true, "", nil},
		// int32
		{int32(6), false, true, "", nil},
		// int64
		{int64(6), false, true, "", nil},
		// uint
		{uint(6), false, true, "", nil},
		// uint8
		{uint8(6), false, true, "", nil},
		// uint16
		{uint16(6), false, true, "", nil},
		// uint32
		{uint32(6), false, true, "", nil},
		// uint64
		{uint64(6), false, true, "", nil},
		// float32
		{float32(6), false, true, "", nil},
		// float64
		{float64(6), false, true, "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, false, false, "unable to convert []int to bool: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, false, false, "unable to convert []uint8 to bool: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, false, false, "unable to convert []int32 to bool: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, false, false, "unable to convert []string to bool: no recipe", nil},
		// struct
		{struct{}{}, false, false, "unable to convert struct {} to bool: struct {} has no Bool() function", nil},

		{SomeStructWithBoolFunc{}, false, true, "", nil},
		{&SomeStructWithBoolFuncPtr{}, false, true, "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
