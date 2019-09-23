package convert

import "testing"

func TestConvertToInt8(t *testing.T) {
	tests := []testCase{
		// nil
		{nil, int8(0), nil, `unable to convert nil to int8: source cannot be nil`, nil},
		// string
		{"6", int8(0), int8(6), "", nil},
		{"Hello World", int8(0), nil, `unable to convert string to int8: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, int8(0), int8(1), "", nil},
		{false, int8(0), int8(0), "", nil},
		// int
		{6, int8(0), int8(6), "", nil},
		// int8
		{int8(6), int8(0), int8(6), "", nil},
		// int16
		{int16(6), int8(0), int8(6), "", nil},
		// uint
		{uint(6), int8(0), int8(6), "", nil},
		// uint16
		{uint16(6), int8(0), int8(6), "", nil},
		// float32
		{float32(6), int8(0), int8(6), "", nil},
		// float64
		{float64(6), int8(0), int8(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, int8(0), nil, "unable to convert []int to int8", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, int8(0), nil, "unable to convert []uint8 to int8", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, int8(0), nil, "unable to convert []int32 to int8", nil},
		{[]string{"H", "e", "l", "l", "o"}, int8(0), nil, "unable to convert []string to int8", nil},
		// struct
		{struct{}{}, int8(0), nil, "unable to convert struct{} to int8", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
