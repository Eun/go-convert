package convert_test

import (
	"testing"

	"github.com/Eun/go-convert/internal/testhelpers"
)

func TestInt64(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, int64(0), int64(0), `unable to convert convert.NilValue to int64: no recipe`, nil},
		// string
		{"6", int64(0), int64(6), "", nil},
		{"", int64(0), int64(0), "", nil},
		{"Hello World", int64(0), int64(0), `unable to convert string to int64: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, int64(0), int64(1), "", nil},
		{false, int64(0), int64(0), "", nil},
		// int
		{6, int64(0), int64(6), "", nil},
		// int8
		{int8(6), int64(0), int64(6), "", nil},
		// int16
		{int16(6), int64(0), int64(6), "", nil},
		// int32
		{int32(6), int64(0), int64(6), "", nil},
		// int64
		{int64(6), int64(0), int64(6), "", nil},
		// uint
		{uint(6), int64(0), int64(6), "", nil},
		// uint8
		{uint8(6), int64(0), int64(6), "", nil},
		// uint16
		{uint16(6), int64(0), int64(6), "", nil},
		// uint32
		{uint32(6), int64(0), int64(6), "", nil},
		// uint64
		{uint64(6), int64(0), int64(6), "", nil},
		// float32
		{float32(6), int64(0), int64(6), "", nil},
		// float64
		{float64(6), int64(0), int64(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, int64(0), int64(0), "unable to convert []int to int64: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, int64(0), int64(0), "unable to convert []uint8 to int64: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, int64(0), int64(0), "unable to convert []int32 to int64: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, int64(0), int64(0), "unable to convert []string to int64: no recipe", nil},
		// struct
		{struct{}{}, int64(0), int64(0), "unable to convert struct {} to int64: no recipe", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
