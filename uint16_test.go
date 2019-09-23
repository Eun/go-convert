package convert

import "testing"

func TestConvertToUint16(t *testing.T) {
	tests := []testCase{
		// nil
		{nil, uint16(0), nil, `unable to convert nil to uint16: source cannot be nil`, nil},
		// string
		{"6", uint16(0), uint16(6), "", nil},
		{"Hello World", uint16(0), nil, `unable to convert string to uint16: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint16(0), uint16(1), "", nil},
		{false, uint16(0), uint16(0), "", nil},
		// int
		{6, uint16(0), uint16(6), "", nil},
		// int16
		{int16(6), uint16(0), uint16(6), "", nil},
		// uint
		{uint(6), uint16(0), uint16(6), "", nil},
		// uint8
		{uint8(6), uint16(0), uint16(6), "", nil},
		// uint16
		{uint16(6), uint16(0), uint16(6), "", nil},
		// float32
		{float32(6), uint16(0), uint16(6), "", nil},
		// float64
		{float64(6), uint16(0), uint16(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint16(0), nil, "unable to convert []int to uint16", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint16(0), nil, "unable to convert []uint8 to uint16", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint16(0), nil, "unable to convert []int32 to uint16", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint16(0), nil, "unable to convert []string to uint16", nil},
		// struct
		{struct{}{}, uint16(0), nil, "unable to convert struct{} to uint16", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
