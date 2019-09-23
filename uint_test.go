package convert

import "testing"

func TestConvertToUint(t *testing.T) {
	tests := []testCase{
		// nil
		{nil, uint(0), nil, `unable to convert nil to uint: source cannot be nil`, nil},
		// string
		{"6", uint(0), uint(6), "", nil},
		{"Hello World", uint(0), nil, `unable to convert string to uint: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint(0), uint(1), "", nil},
		{false, uint(0), uint(0), "", nil},
		// int
		{6, uint(0), uint(6), "", nil},
		// int16
		{int16(6), uint(0), uint(6), "", nil},
		// uint
		{uint(6), uint(0), uint(6), "", nil},
		// uint16
		{uint16(6), uint(0), uint(6), "", nil},
		// float32
		{float32(6), uint(0), uint(6), "", nil},
		// float64
		{float64(6), uint(0), uint(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint(0), nil, "unable to convert []int to uint", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint(0), nil, "unable to convert []uint8 to uint", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint(0), nil, "unable to convert []int32 to uint", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint(0), nil, "unable to convert []string to uint", nil},
		// struct
		{struct{}{}, uint(0), nil, "unable to convert struct{} to uint", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
