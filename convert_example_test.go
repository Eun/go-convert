package convert_test

import (
	"fmt"

	"strings"

	"github.com/Eun/go-convert"
)

func ExampleConvert() {
	// Convert a map of strings to a struct

	// this is the struct we want to convert to
	type User struct {
		ID   int
		Name string
	}

	// this is the data we want to convert
	data := map[string]string{
		"id":   "10",
		"Name": "Joe",
	}

	var user User
	// do the conversion
	convert.MustConvert(data, &user)

	fmt.Printf("Hello %s, your ID is %d\n", user.Name, user.ID)
}

func ExampleNew() {
	type Roles struct {
		IsAdmin     bool
		IsDeveloper bool
	}

	type User struct {
		ID    int
		Name  string
		Roles Roles
	}

	// this is the data we want to convert
	data := map[string]string{
		"id":    "10",
		"Name":  "Joe",
		"roles": "AD", // this user is Admin (A) and Developer (D)
	}

	// create a converter
	conv := convert.New(convert.Options{
		SkipUnknownFields: false,
		Recipes: convert.MustMakeRecipes(
			// convert string into Roles
			func(_ convert.Converter, in string, out *Roles) error {
				out.IsAdmin = false
				out.IsDeveloper = false
				if strings.Contains(in, "A") {
					out.IsAdmin = true
				}
				if strings.Contains(in, "D") {
					out.IsDeveloper = true
				}
				return nil
			},
		),
	})

	var user User
	conv.MustConvert(data, &user)
	// user is now an instance of User
	fmt.Printf("%#v\n", user)
}
