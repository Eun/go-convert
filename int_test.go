package convert

import "testing"

func TestConvertToInt(t *testing.T) {
	tests := []testCase{
		// string
		{"6", 0, 6, "", nil},
		{"Hello World", 0, nil, `unable to convert string to int: strconv.ParseInt: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, 0, 1, "", nil},
		{false, 0, 0, "", nil},
		// int
		{6, 0, 6, "", nil},
		// int16
		{int16(6), 0, 6, "", nil},
		// uint
		{uint(6), 0, 6, "", nil},
		// uint16
		{uint16(6), 0, 6, "", nil},
		// float32
		{float32(6), 0, 6, "", nil},
		// float64
		{float64(6), 0, 6, "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, 0, nil, "unable to convert []int to int", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, 0, nil, "unable to convert []uint8 to int", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, 0, nil, "unable to convert []int32 to int", nil},
		{[]string{"H", "e", "l", "l", "o"}, 0, nil, "unable to convert []string to int", nil},
		// struct
		{struct{}{}, 0, nil, "unable to convert struct{} to int", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
