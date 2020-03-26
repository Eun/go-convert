package convert_test

import (
	"testing"

	"strconv"

	"time"

	"reflect"

	"github.com/Eun/go-convert"
	"github.com/stretchr/testify/require"
)

func TestEdgeCases(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		err := convert.Convert(0, nil)
		require.EqualError(t, err, `destination type cannot be nil`)
	})
	t.Run("interface source", func(t *testing.T) {
		var src interface{}
		var dst int
		err := convert.Convert(src, &dst)
		require.EqualError(t, err, `unable to convert convert.NilValue to int: no recipe`)
	})
	t.Run("nil source", func(t *testing.T) {
		var dst int
		err := convert.Convert(nil, &dst)
		require.EqualError(t, err, `unable to convert convert.NilValue to int: no recipe`)
	})

	t.Run("interface destination", func(t *testing.T) {
		t.Run("string interface", func(t *testing.T) {
			var dst interface{} = ""

			err := convert.Convert("Hello World", &dst)
			require.NoError(t, err)
			require.Equal(t, "Hello World", dst)
		})
		t.Run("nil interface", func(t *testing.T) {
			var dst interface{}

			err := convert.Convert("Hello World", &dst)
			require.NoError(t, err)
			require.Equal(t, "Hello World", dst)
		})
	})
	t.Run("to ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var dst *string
			err := convert.Convert("Hello World", &dst)
			require.NoError(t, err)
			require.Equal(t, "Hello World", *dst)
		})

		t.Run("preset value", func(t *testing.T) {
			d := "Good Bye"
			dst := &d
			err := convert.Convert("Hello World", &dst)
			require.NoError(t, err)
			require.Equal(t, "Hello World", *dst)
		})
	})
}

func TestSimpleString(t *testing.T) {
	var s string
	convert.MustConvert(123, &s)
	require.Equal(t, "123", s)
}

func TestAddNewRecipe(t *testing.T) {
	var s string
	require.NoError(t, convert.Convert(10, &s, convert.Options{
		SkipUnknownFields: false,
		Recipes: []convert.Recipe{
			convert.MustMakeRecipe(func(_ convert.Converter, in int, out *string) error {
				*out = strconv.FormatInt(int64(in), 16)
				return nil
			}),
		},
	}))
	require.Equal(t, "a", s)
}

func TestAddNewRecipeForGeneric(t *testing.T) {
	var s string

	require.NoError(t, convert.Convert(time.Time{}, &s, convert.Options{
		SkipUnknownFields: false,
		Recipes: []convert.Recipe{
			{
				From: convert.StructType,
				To:   reflect.TypeOf(&s),
				Func: func(c convert.Converter, in reflect.Value, out reflect.Value) error {
					require.Fail(t, "Should not be called")
					return nil
				},
			},
		},
	}))
	require.Equal(t, "0001-01-01 00:00:00 +0000 UTC", s)

	require.NoError(t, convert.Convert(struct {
		Name string
	}{
		Name: "Joe",
	}, &s, convert.Options{
		SkipUnknownFields: false,
		Recipes: []convert.Recipe{
			{
				From: convert.StructType,
				To:   reflect.TypeOf(&s),
				Func: func(c convert.Converter, in reflect.Value, out reflect.Value) error {
					out.Elem().Set(reflect.ValueOf("I got you"))
					return nil
				},
			},
		},
	}))
	require.Equal(t, "I got you", s)
}
