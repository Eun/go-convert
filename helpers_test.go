package convert

import (
	"testing"

	"fmt"

	"strings"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Src       interface{}
	DstType   interface{}
	Expected  interface{}
	ErrorText string
	Options   []Option
}

func getTestName(test testCase, index ...int) string {
	var sb strings.Builder
	if len(index) > 0 {
		fmt.Fprintf(&sb, "%d: ", index)
	}
	fmt.Fprintf(&sb, "%s to %s", GetHumanName(test.Src), GetHumanName(test.DstType))
	return sb.String()
}

func runTest(test testCase) func(t *testing.T) {
	return func(t *testing.T) {
		oldDst := test.DstType
		result, err := Convert(test.Src, test.DstType, test.Options...)
		if test.ErrorText == "" {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, test.ErrorText)
		}
		require.Equal(t, test.Expected, result)
		require.Equal(t, oldDst, test.DstType, "dst type changed")
	}
}
