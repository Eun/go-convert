package convert

import "testing"

func TestConvertToInt64(t *testing.T) {
	tests := []testCase{
		// nil
		{nil, int64(0), nil, `unable to convert nil to int64: source cannot be nil`, nil},
		// string
		{"6", int64(0), int64(6), "", nil},
		{"Hello World", int64(0), nil, `unable to convert string to int64: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, int64(0), int64(1), "", nil},
		{false, int64(0), int64(0), "", nil},
		// int
		{6, int64(0), int64(6), "", nil},
		// int16
		{int16(6), int64(0), int64(6), "", nil},
		// int64
		{int64(6), int64(0), int64(6), "", nil},
		// uint
		{uint(6), int64(0), int64(6), "", nil},
		// uint16
		{uint16(6), int64(0), int64(6), "", nil},
		// float32
		{float32(6), int64(0), int64(6), "", nil},
		// float64
		{float64(6), int64(0), int64(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, int64(0), nil, "unable to convert []int to int64", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, int64(0), nil, "unable to convert []uint8 to int64", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, int64(0), nil, "unable to convert []int32 to int64", nil},
		{[]string{"H", "e", "l", "l", "o"}, int64(0), nil, "unable to convert []string to int64", nil},
		// struct
		{struct{}{}, int64(0), nil, "unable to convert struct{} to int64", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
