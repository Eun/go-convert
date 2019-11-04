package testhelpers

import (
	"testing"

	"fmt"

	"strings"

	"reflect"

	convert "github.com/Eun/go-convert"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	Src       interface{}
	Dst       interface{}
	Expected  interface{}
	ErrorText string
	Options   *convert.Options
}

func getTestName(test TestCase, index ...int) string {
	var sb strings.Builder
	if len(index) > 0 {
		fmt.Fprintf(&sb, "%d: ", index)
	}
	src := "nil"
	if test.Src != nil {
		src = reflect.TypeOf(test.Src).String()
	}
	dst := "nil"
	if test.Dst != nil {
		dst = reflect.TypeOf(test.Dst).String()
	}
	fmt.Fprintf(&sb, "%s to %s", src, dst)
	return sb.String()
}

func RunTest(t *testing.T, test TestCase, index ...int) {
	t.Run(getTestName(test, index...), func(t *testing.T) {
		dst := reflect.ValueOf(test.Dst)
		ptrDst := reflect.New(dst.Type())
		ptrDst.Elem().Set(dst)
		var opts []convert.Options
		if test.Options != nil {
			opts = []convert.Options{*test.Options}
		}
		err := convert.Convert(test.Src, ptrDst.Interface(), opts...)
		if test.ErrorText == "" {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, test.ErrorText)
		}
		require.Equal(t, test.Expected, ptrDst.Elem().Interface())
	})
}

func Ptr(v interface{}) interface{} {
	return &v
}

func PtrInt(i int) *int {
	return &i
}

func PtrString(v string) *string {
	return &v
}

func PtrPtrString(v string) **string {
	p := PtrString(v)
	return &p
}
