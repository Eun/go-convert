package convert_test

import (
	"fmt"
	"time"

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
	// Convert a map of strings to a struct
	type User struct {
		ID        int
		Name      string
		CreatedOn time.Time
	}

	// this is the data we want to convert
	data := map[string]string{
		"id":        "10",
		"Name":      "Joe",
		"createdOn": "2001-01-01T00:00:00Z",
	}

	// create a converter
	conv := convert.New(convert.Options{
		SkipUnknownFields: false,
		Recipes: convert.MustMakeRecipes(
			// convert string into time
			func(_ convert.Converter, in string, out *time.Time) error {
				var err error
				*out, err = time.Parse(time.RFC3339, in)
				return err
			},
		),
	})

	var user User
	conv.MustConvert(data, &user)
	// user is now an instance of User
	fmt.Printf("%#v\n", user)
}
