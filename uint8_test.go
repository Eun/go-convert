package convert

import "testing"

func TestConvertToUint8(t *testing.T) {
	tests := []testCase{
		// nil
		{nil, uint8(0), nil, `unable to convert nil to uint8: source cannot be nil`, nil},
		// string
		{"6", uint8(0), uint8(6), "", nil},
		{"Hello World", uint8(0), nil, `unable to convert string to uint8: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint8(0), uint8(1), "", nil},
		{false, uint8(0), uint8(0), "", nil},
		// int
		{6, uint8(0), uint8(6), "", nil},
		// int16
		{int16(6), uint8(0), uint8(6), "", nil},
		// uint
		{uint(6), uint8(0), uint8(6), "", nil},
		// int8
		{uint8(6), uint8(0), uint8(6), "", nil},
		// uint16
		{uint16(6), uint8(0), uint8(6), "", nil},
		// float32
		{float32(6), uint8(0), uint8(6), "", nil},
		// float64
		{float64(6), uint8(0), uint8(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint8(0), nil, "unable to convert []int to uint8", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint8(0), nil, "unable to convert []uint8 to uint8", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint8(0), nil, "unable to convert []int32 to uint8", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint8(0), nil, "unable to convert []string to uint8", nil},
		// struct
		{struct{}{}, uint8(0), nil, "unable to convert struct{} to uint8", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
