package convert

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	tests := []testCase{
		{0, nil, nil, `destination type cannot be nil`, nil},
		//

		// custom converters
		{time.Time{}.AddDate(2000, 0, 0), "", "2001-01-01T00:00:00Z", "", []Option{Options.CustomConverter(func(from time.Time) (string, error) {
			return from.Format(time.RFC3339), nil
		})}},
		{"2001-01-01T00:00:00Z", time.Time{}, time.Time{}.AddDate(2000, 0, 0), "", []Option{Options.CustomConverter(func(from string) (time.Time, error) {
			return time.Parse(time.RFC3339, from)
		})}},

		// unknown
		{"2001-01-01T00:00:00Z", time.Time{}, nil, "unable to convert string to time.Time", nil},
	}

	for i, test := range tests {
		t.Run(getTestName(test, i), runTest(test))
	}
}

func TestGetHumanName(t *testing.T) {
	makeInterface := func(v interface{}) *interface{} {
		return &v
	}
	tests := []struct {
		In       interface{}
		Expected string
	}{
		{nil, "nil"},
		{"", "string"},
		{time.Time{}, "time.Time"},
		{&time.Time{}, "*time.Time"},
		{map[string]interface{}{}, "map[string]interface{}"},
		{&map[string]interface{}{}, "*map[string]interface{}"},
		{&[]string{}, "*[]string"},
		{struct{}{}, "struct{}"},
		{makeInterface(""), "*interface{}(string)"},
	}

	for _, test := range tests {
		require.Equal(t, test.Expected, GetHumanName(test.In))
	}
}
