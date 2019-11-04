package convert_test

import (
	"testing"

	"github.com/Eun/go-convert/internal/testhelpers"
)

func TestSlice(t *testing.T) {
	tests := []testhelpers.TestCase{
		{nil, []int{}, []int{}, "", nil},
		{6, []int{}, []int{}, "unable to convert int to []int: no recipe", nil},
		{[]interface{}{[]int{}}, []int{}, []int{}, "unable to convert []interface {} to []int: unable to convert []int to int: no recipe", nil},
		//
		{"Hello World", []byte{}, []byte("Hello World"), "", nil},
		{"Hello", []int{}, []int{'H', 'e', 'l', 'l', 'o'}, "", nil},
		{"Hello", []string{}, []string{}, "unable to convert string to []string: no recipe", nil},
		// Slice to Slice
		{[]int{1, 2, 3}, []uint{}, []uint{1, 2, 3}, "", nil},
		{[]int{1, 2, 3}, []string{}, []string{"1", "2", "3"}, "", nil},
		{[]int{1, 2, 3}, []interface{}{"", 0.0, 0}, []interface{}{"1", 2.0, 3}, "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
