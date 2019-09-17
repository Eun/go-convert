package convert

import (
	"testing"
)

func TestConvertToSlice(t *testing.T) {
	tests := []testCase{
		{6, []int{}, nil, "unable to convert int to []int", nil},
		{[]interface{}{[]int{}}, []int{}, nil, "unable to convert []interface{} to []int: unable to convert interface{}([]int) to int", nil},

		{"Hello World", []byte{}, []byte("Hello World"), "", nil},
		{"Hello", []int{}, []int{'H', 'e', 'l', 'l', 'o'}, "", nil},
		{"Hello", []string{}, nil, "unable to convert string to []string", nil},
		// Slice to Slice
		{[]int{1, 2, 3}, []uint{}, []uint{1, 2, 3}, "", nil},
		{[]int{1, 2, 3}, []string{}, []string{"1", "2", "3"}, "", nil},
		{[]int{1, 2, 3}, []interface{}{"", 0.0, 0}, []interface{}{"1", 2.0, 3}, "", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
