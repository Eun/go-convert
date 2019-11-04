package convert_test

import (
	"testing"

	"github.com/Eun/go-convert/internal/testhelpers"
)

func TestFloat32(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, float32(0), float32(0), `unable to convert convert.NilValue to float32: no recipe`, nil},
		// string
		{"3.2", float32(0), float32(3.2), "", nil},
		{"Hello World", float32(0), float32(0), `unable to convert string to float32: strconv.ParseFloat: parsing "Hello World": invalid syntax`, nil},
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
		{struct{}{}, float32(0), float32(0), "unable to convert struct {} to float32: no recipe", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
