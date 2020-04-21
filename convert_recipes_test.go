package convert_test

import (
	"testing"

	"github.com/Eun/go-convert"
	"github.com/Eun/go-convert/internal/testhelpers"
)

type StructWithConvertRecipesFunc struct {
	str string
}

func (StructWithConvertRecipesFunc) ConvertRecipes() []convert.Recipe {
	return []convert.Recipe{
		convert.MustMakeRecipe(func(_ convert.Converter, in StructWithConvertRecipesFunc, out *string) error {
			*out = in.str
			return nil
		}),
		convert.MustMakeRecipe(func(_ convert.Converter, in string, out *StructWithConvertRecipesFunc) error {
			out.str = in
			return nil
		}),
	}
}

type StructWithConvertRecipesFuncPtr struct {
	str string
}

func (*StructWithConvertRecipesFuncPtr) ConvertRecipes() []convert.Recipe {
	return []convert.Recipe{
		convert.MustMakeRecipe(func(_ convert.Converter, in StructWithConvertRecipesFuncPtr, out *string) error {
			*out = in.str
			return nil
		}),
		convert.MustMakeRecipe(func(_ convert.Converter, in string, out *StructWithConvertRecipesFuncPtr) error {
			out.str = in
			return nil
		}),
	}
}

func TestConvertRecipes(t *testing.T) {
	tests := []testhelpers.TestCase{
		{StructWithConvertRecipesFunc{"Hello World"}, "", "Hello World", "", nil},
		{&StructWithConvertRecipesFunc{"Hello World"}, "", "Hello World", "", nil},
		{StructWithConvertRecipesFuncPtr{"Hello World"}, "", "Hello World", "", nil},
		{&StructWithConvertRecipesFuncPtr{"Hello World"}, "", "Hello World", "", nil},

		{"Hello World", StructWithConvertRecipesFunc{}, StructWithConvertRecipesFunc{"Hello World"}, "", nil},
		{"Hello World", &StructWithConvertRecipesFunc{}, &StructWithConvertRecipesFunc{"Hello World"}, "", nil},
		{"Hello World", StructWithConvertRecipesFuncPtr{}, StructWithConvertRecipesFuncPtr{"Hello World"}, "", nil},
		{"Hello World", &StructWithConvertRecipesFuncPtr{}, &StructWithConvertRecipesFuncPtr{"Hello World"}, "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
