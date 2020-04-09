package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

func TestFloat64(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, float64(0), float64(0), "", nil},
		// string
		{"3.2", float64(0), float64(3.2), "", nil},
		{"Hello World", float64(0), float64(0), `unable to convert string to float64: strconv.ParseFloat: parsing "Hello World": invalid syntax`, nil},
		{"", float64(0), float64(0), "", nil},
		// bool
		{true, float64(0), float64(1), "", nil},
		{false, float64(0), float64(0), "", nil},
		// int
		{6, float64(0), float64(6), "", nil},
		// int8
		{int8(6), float64(0), float64(6), "", nil},
		// int16
		{int16(6), float64(0), float64(6), "", nil},
		// int32
		{int32(6), float64(0), float64(6), "", nil},
		// int64
		{int64(6), float64(0), float64(6), "", nil},
		// uint
		{uint(6), float64(0), float64(6), "", nil},
		// uint8
		{uint8(6), float64(0), float64(6), "", nil},
		// uint16
		{uint16(6), float64(0), float64(6), "", nil},
		// uint32
		{uint32(6), float64(0), float64(6), "", nil},
		// uint64
		{uint64(6), float64(0), float64(6), "", nil},
		// float32
		{float32(6), float64(0), float64(6), "", nil},
		// float64
		{float64(6), float64(0), float64(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, float64(0), float64(0), "unable to convert []int to float64: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, float64(0), float64(0), "unable to convert []uint8 to float64: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, float64(0), float64(0), "unable to convert []int32 to float64: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, float64(0), float64(0), "unable to convert []string to float64: no recipe", nil},
		// struct
		{struct{}{}, float64(0), float64(0), "unable to convert struct {} to float64: no recipe", nil},
		// time
		{time.Unix(10, 10), float64(10.00000001), float64(10.00000001), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
