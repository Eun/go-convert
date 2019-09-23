package convert

import "testing"

func TestConvertToMap(t *testing.T) {
	type User struct {
		Name string
	}

	type Company struct {
		Name string
	}

	type UserAndCompany struct {
		Name string
		Company
	}

	type privateCompany struct {
		Name string
	}

	type UserAndPrivateCompany struct {
		Name string
		privateCompany
	}

	ptrInt := func(i int) *int {
		return &i
	}

	tests := []testCase{
		// string
		{"Hello World", map[string]interface{}{}, nil, "unable to convert string to map[string]interface{}", nil},

		// map
		{map[string]interface{}{"Foo": true}, map[string]string{}, map[string]string{"Foo": "true"}, "", nil},
		{map[string]string{"Foo": "Bar"}, map[string]interface{}{}, map[string]interface{}{"Foo": "Bar"}, "", nil},

		// respect nested types
		{
			map[string]interface{}{"Int": "3", "Float": "4", "String": "5", "PtrInt": "6", "Slice": []interface{}{"1", "2", "3"}},
			map[string]interface{}{"Int": 0, "Float": 0.0, "PtrInt": ptrInt(0), "Slice": []interface{}{"0", 0, 0.0}},
			map[string]interface{}{"Int": 3, "Float": 4.0, "String": "5", "PtrInt": ptrInt(6), "Slice": []interface{}{"1", 2, 3.0}},
			"",
			nil,
		},

		// respect different key type
		{
			map[int]interface{}{1: 123},
			map[string]interface{}{"1": "String"},
			map[string]interface{}{"1": "123"},
			"",
			nil,
		},

		// respect nested slice
		{
			map[string][]interface{}{"Slice": []interface{}{"1", "2", "3"}},
			map[string][]interface{}{"Slice": {"0", 0, 0.0}},
			map[string][]interface{}{"Slice": {"1", 2, 3.0}},
			"",
			nil,
		},

		// another key type than destination
		{map[string]string{"1": "3"}, map[int]interface{}{1: 0}, map[int]interface{}{1: 3}, "", nil},

		{map[string]string{"1": "3"}, map[interface{}]interface{}{1: 0}, map[interface{}]interface{}{"1": "3"}, "", nil},

		// struct
		{User{"Joe"}, map[string]string{}, map[string]string{"Name": "Joe"}, "", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, map[string]interface{}{}, map[string]interface{}{"Name": "Joe", "Company": Company{"Wood Inc"}}, "", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, map[string]interface{}{"Company": map[string]interface{}{}}, map[string]interface{}{"Name": "Joe", "Company": map[string]interface{}{"Name": "Wood Inc"}}, "", nil},
		{UserAndPrivateCompany{"Joe", privateCompany{"Wood Inc"}}, map[string]interface{}{"privateCompany": map[string]interface{}{}}, map[string]interface{}{"Name": "Joe", "privateCompany": map[string]interface{}{"Name": "Wood Inc"}}, "", nil},

		{map[User]string{User{"Joe"}: "Bar"}, map[string]interface{}{}, nil, "unable to convert map[struct]string to map[string]interface{}: unable to convert convert.User to string", nil},
		{map[string]User{"Foo": User{"Joe"}}, map[string]string{}, nil, "unable to convert map[string]convert.User to map[string]string: unable to convert convert.User to string", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, map[string]string{}, nil, "unable to convert convert.UserAndCompany to map[string]string: unable to convert convert.Company to string", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
