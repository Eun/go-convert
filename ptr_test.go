package convert_test

import (
	"testing"

	"github.com/Eun/go-convert/internal/testhelpers"
)

func TestPtr(t *testing.T) {
	type User struct {
		Name *string
	}

	var nilString *string

	tests := []testhelpers.TestCase{
		// string
		{testhelpers.Ptr("Hello World"), "", "Hello World", "", nil},
		{"Hello World", testhelpers.Ptr(""), testhelpers.Ptr("Hello World"), "", nil},
		{"Hello World", testhelpers.PtrString(""), testhelpers.PtrString("Hello World"), "", nil},
		{testhelpers.Ptr("Hello World"), testhelpers.Ptr(""), testhelpers.Ptr("Hello World"), "", nil},
		{testhelpers.PtrString("Hello World"), "", "Hello World", "", nil},
		{"Hello World", testhelpers.PtrString(""), testhelpers.PtrString("Hello World"), "", nil},
		{testhelpers.PtrString("Hello World"), testhelpers.PtrString(""), testhelpers.PtrString("Hello World"), "", nil},
		//  map
		{testhelpers.Ptr(map[string]interface{}{"Foo": true}), testhelpers.Ptr(map[string]string{}), testhelpers.Ptr(map[string]string{"Foo": "true"}), "", nil},
		{map[string]*string{"Foo": testhelpers.PtrString("Bar")}, map[string]string{}, map[string]string{"Foo": "Bar"}, "", nil},

		// double ptr
		{testhelpers.Ptr(testhelpers.Ptr("Hello World")), testhelpers.Ptr(testhelpers.Ptr("")), testhelpers.Ptr(testhelpers.Ptr("Hello World")), "", nil},
		{testhelpers.Ptr(testhelpers.PtrString("Hello World")), testhelpers.Ptr(testhelpers.PtrString("")), testhelpers.Ptr(testhelpers.PtrString("Hello World")), "", nil},
		{testhelpers.PtrPtrString("Hello World"), testhelpers.PtrPtrString(""), testhelpers.PtrPtrString("Hello World"), "", nil},

		{map[string]string{"Name": "Foo"}, User{}, User{Name: testhelpers.PtrString("Foo")}, "", nil},
		{map[string]string{"Name": "Foo"}, User{Name: testhelpers.PtrString("Bar")}, User{Name: testhelpers.PtrString("Foo")}, "", nil},

		{nilString, 0, 0, "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
