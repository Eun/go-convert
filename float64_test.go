package convert

import "testing"

func TestConvertToFloat64(t *testing.T) {
	tests := []testCase{
		// nil
		{nil, float64(0), nil, `unable to convert nil to float64: source cannot be nil`, nil},
		// string
		{"3.2", float64(0), float64(3.2), "", nil},
		{"Hello World", float64(0), nil, `unable to convert string to float64: strconv.ParseFloat: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, float64(0), float64(1), "", nil},
		{false, float64(0), float64(0), "", nil},
		// int
		{6, float64(0), float64(6), "", nil},
		// int8
		{int8(6), float64(0), float64(6), "", nil},
		// int16
		{int16(6), float64(0), float64(6), "", nil},
		// uint
		{uint(6), float64(0), float64(6), "", nil},
		// uint16
		{uint16(6), float64(0), float64(6), "", nil},
		// float32
		{float32(6), float64(0), float64(6), "", nil},
		// float64
		{float64(6), float64(0), float64(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, float64(0), nil, "unable to convert []int to float64", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, float64(0), nil, "unable to convert []uint8 to float64", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, float64(0), nil, "unable to convert []int32 to float64", nil},
		{[]string{"H", "e", "l", "l", "o"}, float64(0), nil, "unable to convert []string to float64", nil},
		// struct
		{struct{}{}, float64(0), nil, "unable to convert struct{} to float64", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
