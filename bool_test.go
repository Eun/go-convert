package convert

import "testing"

func TestConvertToBool(t *testing.T) {
	tests := []testCase{
		// string
		{"true", false, true, "", nil},
		{"Hello World", false, nil, `unable to convert string to bool: strconv.ParseBool: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, false, true, "", nil},
		{false, false, false, "", nil},
		// int
		{6, false, true, "", nil},
		// int8
		{int8(6), false, true, "", nil},
		// int16
		{int16(6), false, true, "", nil},
		// uint
		{uint(6), false, true, "", nil},
		// uint16
		{uint16(6), false, true, "", nil},
		// float32
		{float32(6), false, true, "", nil},
		// float64
		{float64(6), false, true, "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, false, nil, "unable to convert []int to bool", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, false, nil, "unable to convert []uint8 to bool", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, false, nil, "unable to convert []int32 to bool", nil},
		{[]string{"H", "e", "l", "l", "o"}, false, nil, "unable to convert []string to bool", nil},
		// struct
		{struct{}{}, false, nil, "unable to convert struct{} to bool", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
