<$
import "io/ioutil"
func printFile(f string) {
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(buf))
}
$># go-convert [![Travis](https://img.shields.io/travis/Eun/go-convert.svg)](https://travis-ci.org/Eun/go-convert) [![Codecov](https://img.shields.io/codecov/c/github/Eun/go-convert.svg)](https://codecov.io/gh/Eun/go-convert) [![GoDoc](https://godoc.org/github.com/Eun/go-convert?status.svg)](https://godoc.org/github.com/Eun/go-convert) [![go-report](https://goreportcard.com/badge/github.com/Eun/go-convert)](https://goreportcard.com/report/github.com/Eun/go-convert)
Convert a value into another type.

```bash
go get -u github.com/Eun/go-convert
```
## Usage
```go
<$ printFile("./examples/usage/usage.go") $>
```

## Recipe system
_go-convert_ uses a recipe system that defines how and which types should be converted in which type.
A lot of recipes are already builtin (see [recipes.go](recipes.go)), however you can add your own or overwrite the builtin ones.
```go
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

```

## Adding inline recipes
You can also add recipes inline by implementing a `ConvertRecipes() []Recipe` function.  
Example:
```go
type Roles struct {
    IsAdmin     bool
    IsDeveloper bool
}

type User struct {
    ID    int
    Name  string
    Roles Roles
}

func (user *User) ConvertRecipes() []Recipe {
    return []Recipe{
    	convert.MustMakeRecipes(
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
    }
}

// this is the data we want to convert
data := map[string]string{
    "id":    "10",
    "Name":  "Joe",
    "roles": "AD", // this user is Admin (A) and Developer (D)
}

var user User
converter.MustConvert(data, &user)
// user is now an instance of User
fmt.Printf("%#v\n", user)
```


### Notice
This library is using reflection so be aware it might be slow in your usecase.  
