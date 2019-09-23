package convert

import "testing"

func TestConvertToString(t *testing.T) {
	tests := []testCase{
		// nil
		{nil, "", nil, `unable to convert nil to string: source cannot be nil`, nil},
		// string
		{"Hello World", "", "Hello World", "", nil},
		// bool
		{true, "", "true", "", nil},
		// int
		{6, "", "6", "", nil},
		// int16
		{int16(6), "", "6", "", nil},
		// uint
		{uint(6), "", "6", "", nil},
		// uint16
		{uint16(6), "", "6", "", nil},
		// float32
		{float32(6), "", "6.000000", "", nil},
		// float64
		{float64(6), "", "6.000000", "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, "", "Hello", "", nil},
		{[]string{"H", "e", "l", "l", "o"}, "", nil, "unable to convert []string to string", nil},
		// struct
		{struct{}{}, "", nil, "unable to convert struct{} to string", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
