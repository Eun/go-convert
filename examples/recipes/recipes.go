package main

import (
	"fmt"
	"strings"

	"github.com/Eun/go-convert"
)

type Roles struct {
	IsAdmin     bool
	IsDeveloper bool
}

type User struct {
	ID    int
	Name  string
	Roles Roles
}

func main() {
	// this is the data we want to convert
	data := map[string]string{
		"id":    "10",
		"Name":  "Joe",
		"roles": "AD", // this user is Admin (A) and Developer (D)
	}

	// create a converter
	conv := convert.New(convert.Options{
		Recipes: convert.MustMakeRecipes(
			// convert string into Roles
			func(_ convert.Converter, in string, out *Roles) error {
				(*out).IsAdmin = false
				(*out).IsDeveloper = false
				if strings.Contains(in, "A") {
					(*out).IsAdmin = true
				}
				if strings.Contains(in, "D") {
					(*out).IsDeveloper = true
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
