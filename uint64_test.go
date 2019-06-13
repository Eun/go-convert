package convert

import "testing"

func TestConvertToUint64(t *testing.T) {
	tests := []testCase{
		// string
		{"6", uint64(0), uint64(6), "", nil},
		{"Hello World", uint64(0), nil, `unable to convert string to uint64: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint64(0), uint64(1), "", nil},
		{false, uint64(0), uint64(0), "", nil},
		// int
		{6, uint64(0), uint64(6), "", nil},
		// int16
		{int16(6), uint64(0), uint64(6), "", nil},
		// uint
		{uint(6), uint64(0), uint64(6), "", nil},
		// uint16
		{uint16(6), uint64(0), uint64(6), "", nil},
		// uint64
		{uint64(6), uint64(0), uint64(6), "", nil},
		// float32
		{float32(6), uint64(0), uint64(6), "", nil},
		// float64
		{float64(6), uint64(0), uint64(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint64(0), nil, "unable to convert []int to uint64", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint64(0), nil, "unable to convert []uint8 to uint64", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint64(0), nil, "unable to convert []int32 to uint64", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint64(0), nil, "unable to convert []string to uint64", nil},
		// struct
		{struct{}{}, uint64(0), nil, "unable to convert struct{} to uint64", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
