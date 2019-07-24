package denary_test

import (
	"github.com/reiver/go-denary"

	"testing"
)

func TestParseString(t *testing.T) {

	tests := []struct{
		Src string
		Expected denary.Type
	}{
		{
			Src:                 "-128",
			Expected: denary.Int8(-128),
		},
		{
			Src:                  "-128",
			Expected: denary.Int16(-128),
		},
		{
			Src:                  "-128",
			Expected: denary.Int32(-128),
		},
		{
			Src:                  "-128",
			Expected: denary.Int64(-128),
		},



		{
			Src:                 "-1",
			Expected: denary.Int8(-1),
		},
		{
			Src:                  "-1",
			Expected: denary.Int16(-1),
		},
		{
			Src:                  "-1",
			Expected: denary.Int32(-1),
		},
		{
			Src:                  "-1",
			Expected: denary.Int64(-1),
		},



		{
			Src:                 "0",
			Expected: denary.Int8(0),
		},
		{
			Src:                  "0",
			Expected: denary.Int16(0),
		},
		{
			Src:                  "0",
			Expected: denary.Int32(0),
		},
		{
			Src:                  "0",
			Expected: denary.Int64(0),
		},



		{
			Src:                  "0",
			Expected: denary.Uint8(0),
		},
		{
			Src:                   "0",
			Expected: denary.Uint16(0),
		},
		{
			Src:                   "0",
			Expected: denary.Uint32(0),
		},
		{
			Src:                   "0",
			Expected: denary.Uint64(0),
		},



		{
			Src:                 "1",
			Expected: denary.Int8(1),
		},
		{
			Src:                  "1",
			Expected: denary.Int16(1),
		},
		{
			Src:                  "1",
			Expected: denary.Int32(1),
		},
		{
			Src:                  "1",
			Expected: denary.Int64(1),
		},



		{
			Src:                  "1",
			Expected: denary.Uint8(1),
		},
		{
			Src:                   "1",
			Expected: denary.Uint16(1),
		},
		{
			Src:                   "1",
			Expected: denary.Uint32(1),
		},
		{
			Src:                   "1",
			Expected: denary.Uint64(1),
		},



		{
			Src:                 "127",
			Expected: denary.Int8(127),
		},
		{
			Src:                  "127",
			Expected: denary.Int16(127),
		},
		{
			Src:                  "127",
			Expected: denary.Int32(127),
		},
		{
			Src:                  "127",
			Expected: denary.Int64(127),
		},



		{
			Src:                  "127",
			Expected: denary.Uint8(127),
		},
		{
			Src:                   "127",
			Expected: denary.Uint16(127),
		},
		{
			Src:                   "127",
			Expected: denary.Uint32(127),
		},
		{
			Src:                   "127",
			Expected: denary.Uint64(127),
		},









		{
			Src:                  "-5",
			Expected: denary.Int64(-5),
		},
	}

	for testNumber, test := range tests {

		result := denary.Parse(test.Src)

		actual, ok := result.Unwrap()
		if !ok {
			err := result.Validate()
			t.Errorf("For test #%d, could not unwrap the result (%t), because: (%T) %q", testNumber, ok, err, err)
			t.Logf("SRC: «%s»", test.Src)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the denary value that was actually gotten is not what was expected.", testNumber, )
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
