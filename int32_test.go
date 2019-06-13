package convert

import "testing"

func TestConvertToInt32(t *testing.T) {
	tests := []testCase{
		// string
		{"6", int32(0), int32(6), "", nil},
		{"Hello World", int32(0), nil, `unable to convert string to int32: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, int32(0), int32(1), "", nil},
		{false, int32(0), int32(0), "", nil},
		// int
		{6, int32(0), int32(6), "", nil},
		// int16
		{int16(6), int32(0), int32(6), "", nil},
		// int32
		{int32(6), int32(0), int32(6), "", nil},
		// uint
		{uint(6), int32(0), int32(6), "", nil},
		// uint16
		{uint16(6), int32(0), int32(6), "", nil},
		// float32
		{float32(6), int32(0), int32(6), "", nil},
		// float64
		{float64(6), int32(0), int32(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, int32(0), nil, "unable to convert []int to int32", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, int32(0), nil, "unable to convert []uint8 to int32", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, int32(0), nil, "unable to convert []int32 to int32", nil},
		{[]string{"H", "e", "l", "l", "o"}, int32(0), nil, "unable to convert []string to int32", nil},
		// struct
		{struct{}{}, int32(0), nil, "unable to convert struct{} to int32", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
