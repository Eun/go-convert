package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

func TestInt16(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, int16(0), int16(0), `unable to convert convert.NilValue to int16: no recipe`, nil},
		// string
		{"6", int16(0), int16(6), "", nil},
		{"", int16(0), int16(0), "", nil},
		{"Hello World", int16(0), int16(0), `unable to convert string to int16: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, int16(0), int16(1), "", nil},
		{false, int16(0), int16(0), "", nil},
		// int
		{6, int16(0), int16(6), "", nil},
		// int8
		{int8(6), int16(0), int16(6), "", nil},
		// int16
		{int16(6), int16(0), int16(6), "", nil},
		// int32
		{int32(6), int16(0), int16(6), "", nil},
		// int64
		{int64(6), int16(0), int16(6), "", nil},
		// uint
		{uint(6), int16(0), int16(6), "", nil},
		// uint8
		{uint8(6), int16(0), int16(6), "", nil},
		// uint16
		{uint16(6), int16(0), int16(6), "", nil},
		// uint32
		{uint32(6), int16(0), int16(6), "", nil},
		// uint64
		{uint64(6), int16(0), int16(6), "", nil},
		// float32
		{float32(6), int16(0), int16(6), "", nil},
		// float64
		{float64(6), int16(0), int16(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, int16(0), int16(0), "unable to convert []int to int16: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, int16(0), int16(0), "unable to convert []uint8 to int16: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, int16(0), int16(0), "unable to convert []int32 to int16: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, int16(0), int16(0), "unable to convert []string to int16: no recipe", nil},
		// struct
		{struct{}{}, int16(0), int16(0), "unable to convert struct {} to int16: no recipe", nil},
		// time
		{time.Unix(10, 10), int16(10), int16(10), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
