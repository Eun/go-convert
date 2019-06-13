package convert

import "testing"

func TestConvertToInt16(t *testing.T) {
	tests := []testCase{
		// string
		{"6", int16(0), int16(6), "", nil},
		{"Hello World", int16(0), nil, `unable to convert string to int16: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, int16(0), int16(1), "", nil},
		{false, int16(0), int16(0), "", nil},
		// int
		{6, int16(0), int16(6), "", nil},
		// int16
		{int16(6), int16(0), int16(6), "", nil},
		// uint
		{uint(6), int16(0), int16(6), "", nil},
		// uint16
		{uint16(6), int16(0), int16(6), "", nil},
		// float32
		{float32(6), int16(0), int16(6), "", nil},
		// float64
		{float64(6), int16(0), int16(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, int16(0), nil, "unable to convert []int to int16", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, int16(0), nil, "unable to convert []uint8 to int16", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, int16(0), nil, "unable to convert []int32 to int16", nil},
		{[]string{"H", "e", "l", "l", "o"}, int16(0), nil, "unable to convert []string to int16", nil},
		// struct
		{struct{}{}, int16(0), nil, "unable to convert struct{} to int16", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
