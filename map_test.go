package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

func TestMap(t *testing.T) {
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

	type TimeStruct struct {
		CreatedOn *time.Time
	}

	beginningOfTime := time.Unix(0, 0).UTC()

	tests := []testhelpers.TestCase{
		// nil
		{nil, map[string]interface{}{}, map[string]interface{}{}, "", nil},
		// string
		{"Hello World", map[string]interface{}{}, map[string]interface{}{}, "unable to convert string to map[string]interface {}: no recipe", nil},
		// map
		{map[string]interface{}{"Foo": true}, map[string]string{}, map[string]string{"Foo": "true"}, "", nil},
		{map[string]string{"Foo": "Bar"}, map[string]interface{}{}, map[string]interface{}{"Foo": "Bar"}, "", nil},
		//
		// respect nested types
		{
			map[string]interface{}{"Int": "3", "Float": "4", "String": "5", "PtrInt": "6", "Slice": []interface{}{"1", "2", "3"}},
			map[string]interface{}{"Int": 0, "Float": 0.0, "PtrInt": testhelpers.PtrInt(0), "Slice": []interface{}{"0", 0, 0.0}},
			map[string]interface{}{"Int": 3, "Float": 4.0, "String": "5", "PtrInt": testhelpers.PtrInt(6), "Slice": []interface{}{"1", 2, 3.0}},
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
			map[string][]interface{}{"Slice": {"1", "2", "3"}},
			map[string][]interface{}{"Slice": {"0", 0, 0.0}},
			map[string][]interface{}{"Slice": {"1", 2, 3.0}},
			"",
			nil,
		},

		// another key type than destination
		{map[string]string{"1": "3"}, map[int]interface{}{1: 0}, map[int]interface{}{1: 3}, "", nil},
		//
		{map[string]string{"1": "3"}, map[interface{}]interface{}{1: 0}, map[interface{}]interface{}{"1": "3"}, "", nil},
		//
		// null interface
		{map[string]interface{}{"Foo": nil}, map[string][]string{}, map[string][]string{"Foo": nil}, "", nil},
		//
		// struct
		{User{"Joe"}, map[string]string{}, map[string]string{"Name": "Joe"}, "", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, map[string]interface{}{}, map[string]interface{}{"Name": "Joe", "Company": Company{"Wood Inc"}}, "", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, map[string]interface{}{"Company": map[string]interface{}{}}, map[string]interface{}{"Name": "Joe", "Company": map[string]interface{}{"Name": "Wood Inc"}}, "", nil},
		{UserAndPrivateCompany{"Joe", privateCompany{"Wood Inc"}}, map[string]interface{}{"privateCompany": map[string]interface{}{}}, map[string]interface{}{"Name": "Joe", "privateCompany": map[string]interface{}{"Name": "Wood Inc"}}, "", nil},

		{map[User]string{{"Joe"}: "Bar"}, map[string]interface{}{}, map[string]interface{}{}, "unable to convert map[convert_test.User]string to map[string]interface {}: unable to convert convert_test.User to string: convert_test.User has no String() function", nil},
		{map[string]User{"Foo": {"Joe"}}, map[string]string{}, map[string]string{}, "unable to convert map[string]convert_test.User to map[string]string: unable to convert convert_test.User to string: convert_test.User has no String() function", nil},
		{UserAndCompany{"Joe", Company{"Wood Inc"}}, map[string]string{}, map[string]string{}, "unable to convert convert_test.UserAndCompany to map[string]string: unable to convert convert_test.Company to string: convert_test.Company has no String() function", nil},

		{TimeStruct{}, map[string]string{}, map[string]string{"CreatedOn": "0001-01-01 00:00:00 +0000 UTC"}, "", nil},
		{TimeStruct{CreatedOn: &beginningOfTime}, map[string]string{}, map[string]string{"CreatedOn": "1970-01-01 00:00:00 +0000 UTC"}, "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
