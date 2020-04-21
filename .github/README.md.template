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
<$ printFile("./examples/recipes/recipes.go") $>
```

## Adding inline recipes
You can also add recipes inline by implementing a `ConvertRecipes() []Recipe` function.  
Example:
```go
<$ printFile("./examples/inline_recipes/inline_recipes.go") $>
```


### Notice
This library is using reflection so be aware it might be slow in your usecase.  
