package convert

import (
	"fmt"
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

	// do the conversion
	user := MustConvert(data, User{})

	// user is now an instance of User
	fmt.Printf("Hello %s, your ID is %d\n", user.(User).Name, user.(User).ID)
}
