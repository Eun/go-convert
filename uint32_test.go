package convert

import "testing"

func TestConvertToUint32(t *testing.T) {
	tests := []testCase{
		// string
		{"6", uint32(0), uint32(6), "", nil},
		{"Hello World", uint32(0), nil, `unable to convert string to uint32: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint32(0), uint32(1), "", nil},
		{false, uint32(0), uint32(0), "", nil},
		// int
		{6, uint32(0), uint32(6), "", nil},
		// int16
		{int16(6), uint32(0), uint32(6), "", nil},
		// int32
		{int32(6), uint32(0), uint32(6), "", nil},
		// uint
		{uint(6), uint32(0), uint32(6), "", nil},
		// uint16
		{uint16(6), uint32(0), uint32(6), "", nil},
		// uint16
		{uint32(6), uint32(0), uint32(6), "", nil},
		// float32
		{float32(6), uint32(0), uint32(6), "", nil},
		// float64
		{float64(6), uint32(0), uint32(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint32(0), nil, "unable to convert []int to uint32", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint32(0), nil, "unable to convert []uint8 to uint32", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint32(0), nil, "unable to convert []int32 to uint32", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint32(0), nil, "unable to convert []string to uint32", nil},
		// struct
		{struct{}{}, uint32(0), nil, "unable to convert struct{} to uint32", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
