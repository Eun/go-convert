package convert_test

import (
	"testing"

	convert "github.com/Eun/go-convert"
	"github.com/Eun/go-convert/internal/testhelpers"
)

func TestStruct(t *testing.T) {
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

	type NilField struct {
		OptionalField *string
	}

	tests := []testhelpers.TestCase{
		// nil
		{nil, struct{}{}, struct{}{}, "", nil},

		// string
		{"Hello World", struct{}{}, struct{}{}, "unable to convert string to struct {}: no recipe", nil},

		// map
		{map[string]interface{}{"Ok": true}, Foo{}, Foo{true}, "", nil},
		{map[string]string{"Name": "Joe"}, User{}, User{"Joe"}, "", nil},
		{map[string]interface{}{"Name": "Joe", "Company": map[string]interface{}{"Name": "Wood Inc"}}, UserAndCompany{}, UserAndCompany{"Joe", Company{"Wood Inc"}}, "", nil},
		{map[bool]string{true: "Bar"}, Bool{}, Bool{"Bar"}, "", nil},
		{map[string]interface{}{}, NilField{}, NilField{}, "", nil},
		{map[string]interface{}{"OptionalField": "Hello World"}, NilField{}, NilField{testhelpers.PtrString("Hello World")}, "", nil},

		{map[string]interface{}{"Foo": "Bar"}, User{}, User{}, `unable to convert map[string]interface {} to convert_test.User: unable to find Foo in convert_test.User`, nil},
		{map[string]interface{}{"Foo": "Bar"}, User{}, User{}, "", &convert.Options{SkipUnknownFields: true}},

		// should be unable to convert key
		{map[User]string{{}: ""}, struct{}{}, struct{}{}, `unable to convert map[convert_test.User]string to struct {}: unable to convert convert_test.User to string: convert_test.User has no String() function`, nil},
		// should be unable to convert value
		{map[string]User{"Foo": {}}, struct{ Foo Foo }{}, struct{ Foo Foo }{}, `unable to convert map[string]convert_test.User to struct { Foo convert_test.Foo }: unable to convert convert_test.User to convert_test.Foo: unable to find Name in convert_test.Foo`, nil},
		//
		// struct
		{User{"Joe"}, User{}, User{"Joe"}, "", nil},
		{User{"Joe"}, AnotherUser{}, AnotherUser{"Joe"}, "", nil},
		{User{"Joe"}, AnotherByteUser{}, AnotherByteUser{[]byte("Joe")}, "", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, AnotherUserAndCompany{}, AnotherUserAndCompany{"Joe", Company{"Wood Inc"}}, "", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, AnotherUserAndAnotherCompany{}, AnotherUserAndAnotherCompany{"Joe", AnotherCompany{"Wood Inc"}}, "", nil},

		{UserAndCompany{"Joe", Company{"Wood Inc"}}, User{}, User{}, "unable to convert convert_test.UserAndCompany to convert_test.User: unable to find Company in convert_test.User", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, User{}, User{"Joe"}, "", &convert.Options{SkipUnknownFields: true}},

		{UserAndCompany{"Joe", Company{"Wood Inc"}}, AnotherUserAndCompanyString{}, AnotherUserAndCompanyString{}, "unable to convert convert_test.UserAndCompany to convert_test.AnotherUserAndCompanyString: unable to convert convert_test.Company to string: convert_test.Company has no String() function", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
