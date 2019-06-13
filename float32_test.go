package convert

import "testing"

func TestConvertToFloat32(t *testing.T) {
	tests := []testCase{
		// string
		{"3.2", float32(0), float32(3.2), "", nil},
		{"Hello World", float32(0), nil, `unable to convert string to float32: strconv.ParseFloat: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, float32(0), float32(1), "", nil},
		{false, float32(0), float32(0), "", nil},
		// int
		{6, float32(0), float32(6), "", nil},
		// int8
		{int8(6), float32(0), float32(6), "", nil},
		// int16
		{int16(6), float32(0), float32(6), "", nil},
		// uint
		{uint(6), float32(0), float32(6), "", nil},
		// uint16
		{uint16(6), float32(0), float32(6), "", nil},
		// float32
		{float32(6), float32(0), float32(6), "", nil},
		// float64
		{float64(6), float32(0), float32(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, float32(0), nil, "unable to convert []int to float32", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, float32(0), nil, "unable to convert []uint8 to float32", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, float32(0), nil, "unable to convert []int32 to float32", nil},
		{[]string{"H", "e", "l", "l", "o"}, float32(0), nil, "unable to convert []string to float32", nil},
		// struct
		{struct{}{}, float32(0), nil, "unable to convert struct{} to float32", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
