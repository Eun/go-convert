package convert

import (
	"time"

	"fmt"
)

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
	converter := New(
		Options.CustomConverter(func(from string) (time.Time, error) {
			return time.Parse(time.RFC3339, from)
		}),
	)

	user := converter.MustConvert(data, User{})

	// user is now an instance of User
	fmt.Printf("%#v\n", user.(User))
}
