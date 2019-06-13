package convert

import "testing"

func TestConvertPtr(t *testing.T) {
	ptr := func(v interface{}) interface{} {
		return &v
	}
	ptrString := func(v string) *string {
		return &v
	}
	ptrptrString := func(v string) **string {
		p := ptrString(v)
		return &p
	}

	tests := []testCase{
		// string
		{ptr("Hello World"), "", "Hello World", "", nil},
		{"Hello World", ptr(""), ptr("Hello World"), "", nil},
		{ptr("Hello World"), ptr(""), ptr("Hello World"), "", nil},
		{ptrString("Hello World"), "", "Hello World", "", nil},
		{"Hello World", ptrString(""), ptrString("Hello World"), "", nil},
		{ptrString("Hello World"), ptrString(""), ptrString("Hello World"), "", nil},
		{ptrString("Hello World"), ptrString(""), "Hello World", "", []Option{Options.SkipPointers()}},
		// // map
		{ptr(map[string]interface{}{"Foo": true}), ptr(map[string]string{}), ptr(map[string]string{"Foo": "true"}), "", nil},
		{map[string]*string{"Foo": ptrString("Bar")}, map[string]string{}, map[string]string{"Foo": "Bar"}, "", nil},

		// double ptr
		{ptr(ptr("Hello World")), ptr(ptr("")), ptr(ptr("Hello World")), "", nil},
		{ptr(ptrString("Hello World")), ptr(ptrString("")), ptr(ptrString("Hello World")), "", nil},
		{ptrptrString("Hello World"), ptrptrString(""), ptrptrString("Hello World"), "", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
