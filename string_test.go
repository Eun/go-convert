package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

type SomeStructWithStringFunc struct {
}

func (SomeStructWithStringFunc) String() string {
	return "Hello World"
}

type SomeStructWithStringFuncPtr struct {
}

func (*SomeStructWithStringFuncPtr) String() string {
	return "Hello World"
}

func TestString(t *testing.T) {
	beginningOfTime := time.Unix(0, 0).UTC()
	tests := []testhelpers.TestCase{
		// nil
		{nil, "", "", "", nil},
		// string
		{"Hello World", "", "Hello World", "", nil},
		// bool
		{true, "", "true", "", nil},
		// int
		{6, "", "6", "", nil},
		// int8
		{int8(6), "", "6", "", nil},
		// int16
		{int16(6), "", "6", "", nil},
		// int32
		{int32(6), "", "6", "", nil},
		// int64
		{int64(6), "", "6", "", nil},
		// uint
		{uint(6), "", "6", "", nil},
		// uint8
		{uint8(6), "", "6", "", nil},
		// uint16
		{uint16(6), "", "6", "", nil},
		// uint32
		{uint32(6), "", "6", "", nil},
		// uint64
		{uint64(6), "", "6", "", nil},
		// float32
		{float32(6), "", "6.000000", "", nil},
		// float32 with format
		{float32(6), "%06.2f", "006.00", "", nil},
		// float64
		{float64(6), "", "6.000000", "", nil},
		// float64 with format
		{float64(6), "%06.2f", "006.00", "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]int8{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]int16{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]int32{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]int64{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]uint{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]uint8{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]uint16{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]uint32{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]uint64{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]string{"H", "e", "l", "l", "o"}, "", "", "unable to convert []string to string: no recipe", nil},
		// struct
		{struct{}{}, "", "", "unable to convert struct {} to string: struct {} has no String() function", nil},

		{SomeStructWithStringFunc{}, "Hello World", "Hello World", "", nil},
		{&SomeStructWithStringFuncPtr{}, "Hello World", "Hello World", "", nil},
		{beginningOfTime, "", "1970-01-01 00:00:00 +0000 UTC", "", nil},
		{&beginningOfTime, "", "1970-01-01 00:00:00 +0000 UTC", "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
