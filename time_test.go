package convert_test

import (
	"testing"
	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

func TestTime(t *testing.T) {
	someTime := func() time.Time {
		return time.Time{}.AddDate(2005, 1, 0)
	}

	tests := []testhelpers.TestCase{
		// nil
		{nil, time.Time{}, time.Time{}, `unable to convert convert.NilValue to time.Time: no recipe`, nil},
		// string
		{"2/1/2006", time.Time{}, someTime(), "", nil},
		// int
		{100, time.Time{}, time.Unix(100, 0), "", nil},
		// int8
		{int8(100), time.Time{}, time.Unix(100, 0), "", nil},
		// int16
		{int16(100), time.Time{}, time.Unix(100, 0), "", nil},
		// int32
		{int32(100), time.Time{}, time.Unix(100, 0), "", nil},
		// int64
		{int64(100), time.Time{}, time.Unix(100, 0), "", nil},
		// uint
		{uint(100), time.Time{}, time.Unix(100, 0), "", nil},
		// uint8
		{uint8(100), time.Time{}, time.Unix(100, 0), "", nil},
		// uint16
		{uint16(100), time.Time{}, time.Unix(100, 0), "", nil},
		// uint32
		{uint32(100), time.Time{}, time.Unix(100, 0), "", nil},
		// uint64
		{uint64(100), time.Time{}, time.Unix(100, 0), "", nil},
		// float32
		{float32(100), time.Time{}, time.Unix(100, 0), "", nil},
		// float64
		{float64(100), time.Time{}, time.Unix(100, 0), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}

func TestA(t *testing.T) {

}
