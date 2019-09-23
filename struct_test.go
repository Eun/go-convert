package convert

import "testing"

func TestConvertToStruct(t *testing.T) {
	type Foo struct {
		Ok bool
	}

	type Bool struct {
		True string
	}

	type User struct {
		Name string
	}

	type AnotherUser struct {
		Name string
	}

	type AnotherByteUser struct {
		Name []byte
	}

	type Company struct {
		Name string
	}

	type UserAndCompany struct {
		Name string
		Company
	}

	type AnotherUserAndCompany struct {
		Name string
		Company
	}

	type AnotherCompany struct {
		Name string
	}

	type AnotherUserAndAnotherCompany struct {
		Name    string
		Company AnotherCompany
	}

	type AnotherUserAndCompanyString struct {
		Name    string
		Company string
	}

	tests := []testCase{
		// nil
		{nil, struct{}{}, nil, `unable to convert nil to struct{}: source cannot be nil`, nil},

		// string
		{"Hello World", struct{}{}, nil, "unable to convert string to struct{}", nil},

		// map
		{map[string]interface{}{"Ok": true}, Foo{}, Foo{true}, "", nil},
		{map[string]string{"Name": "Joe"}, User{}, User{"Joe"}, "", nil},
		{map[string]interface{}{"Name": "Joe", "Company": map[string]interface{}{"Name": "Wood Inc"}}, UserAndCompany{}, UserAndCompany{"Joe", Company{"Wood Inc"}}, "", nil},
		{map[bool]string{true: "Bar"}, Bool{}, Bool{"Bar"}, "", nil},
		{map[User]string{User{"Joe"}: "Bar"}, User{}, nil, "unable to convert map[struct]string to convert.User: unable to convert convert.User to string", nil},
		{map[string]User{"Name": User{"Joe"}}, User{}, nil, "unable to convert map[string]convert.User to convert.User: unable to convert convert.User to string", nil},

		{map[string]interface{}{"Foo": "Bar"}, User{}, nil, `unable to convert map[string]interface{} to convert.User: unable to find Foo in convert.User`, nil},
		{map[string]interface{}{"Foo": "Bar"}, User{}, User{}, "", []Option{Options.SkipUnknownFields()}},

		// struct
		{User{"Joe"}, User{}, User{"Joe"}, "", nil},
		{User{"Joe"}, AnotherUser{}, AnotherUser{"Joe"}, "", nil},
		{User{"Joe"}, AnotherByteUser{}, AnotherByteUser{[]byte("Joe")}, "", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, AnotherUserAndCompany{}, AnotherUserAndCompany{"Joe", Company{"Wood Inc"}}, "", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, AnotherUserAndAnotherCompany{}, AnotherUserAndAnotherCompany{"Joe", AnotherCompany{"Wood Inc"}}, "", nil},

		{UserAndCompany{"Joe", Company{"Wood Inc"}}, User{}, nil, "unable to convert convert.UserAndCompany to convert.User: unable to find Company in convert.User", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, User{}, User{"Joe"}, "", []Option{Options.SkipUnknownFields()}},

		{UserAndCompany{"Joe", Company{"Wood Inc"}}, AnotherUserAndCompanyString{}, nil, "unable to convert convert.UserAndCompany to convert.AnotherUserAndCompanyString: unable to convert convert.Company to string", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}
