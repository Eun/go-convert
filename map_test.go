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
		{map[string]string{"Foo": "3", "Bar": "4", "Beef": "5"}, map[string]interface{}{"Foo": 3, "Bar": 4.0}, map[string]interface{}{"Foo": 3, "Bar": 4.0, "Beef": "5"}, "", nil},
		{map[string]string{"Foo": "3"}, map[string]interface{}{"Foo": ptrInt(3)}, map[string]interface{}{"Foo": ptrInt(3)}, "", nil},
		{map[string]string{"1": "3"}, map[int]interface{}{1: 0}, map[int]interface{}{1: 3}, "", nil},

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
