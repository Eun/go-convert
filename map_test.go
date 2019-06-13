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

	tests := []testCase{
		// string
		{"Hello World", map[string]interface{}{}, nil, "unable to convert string to map[string]interface{}", nil},

		// map
		{map[string]interface{}{"Foo": true}, map[string]string{}, map[string]string{"Foo": "true"}, "", nil},
		{map[string]string{"Foo": "Bar"}, map[string]interface{}{}, map[string]interface{}{"Foo": "Bar"}, "", nil},

		// // struct
		{User{"Joe"}, map[string]string{}, map[string]string{"Name": "Joe"}, "", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, map[string]interface{}{}, map[string]interface{}{"Name": "Joe", "Company": Company{"Wood Inc"}}, "", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, map[string]interface{}{}, map[string]interface{}{"Name": "Joe", "Company": map[string]interface{}{"Name": "Wood Inc"}}, "", []Option{Options.ConvertEmbeddedStructToParentType()}},

		{map[User]string{User{"Joe"}: "Bar"}, map[string]interface{}{}, nil, "unable to convert map[struct]string to map[string]interface{}: unable to convert convert.User to string", nil},
		{map[string]User{"Foo": User{"Joe"}}, map[string]string{}, nil, "unable to convert map[string]convert.User to map[string]string: unable to convert convert.User to string", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, map[string]string{}, nil, "unable to convert convert.UserAndCompany to map[string]string: unable to convert convert.Company to string", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
