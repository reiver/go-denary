package denary_test

import (
	"github.com/reiver/go-denary"

	"testing"
)

func TestParseType(t *testing.T) {

	tests := []struct{
		Src denary.Type
		Expected denary.Type
	}{
		{
			Src:      denary.Int8(-128),
			Expected: denary.Int8(-128),
		},
		{
			Src:      denary.Int16(-128),
			Expected: denary.Int16(-128),
		},
		{
			Src:      denary.Int32(-128),
			Expected: denary.Int32(-128),
		},
		{
			Src:      denary.Int64(-128),
			Expected: denary.Int64(-128),
		},



		{
			Src:      denary.Int8(-1),
			Expected: denary.Int8(-1),
		},
		{
			Src:      denary.Int16(-1),
			Expected: denary.Int16(-1),
		},
		{
			Src:      denary.Int32(-1),
			Expected: denary.Int32(-1),
		},
		{
			Src:      denary.Int64(-1),
			Expected: denary.Int64(-1),
		},



		{
			Src:      denary.Int8(0),
			Expected: denary.Int8(0),
		},
		{
			Src:      denary.Int16(0),
			Expected: denary.Int16(0),
		},
		{
			Src:      denary.Int32(0),
			Expected: denary.Int32(0),
		},
		{
			Src:      denary.Int64(0),
			Expected: denary.Int64(0),
		},



		{
			Src:      denary.Uint8(0),
			Expected: denary.Uint8(0),
		},
		{
			Src:      denary.Uint16(0),
			Expected: denary.Uint16(0),
		},
		{
			Src:      denary.Uint32(0),
			Expected: denary.Uint32(0),
		},
		{
			Src:      denary.Uint64(0),
			Expected: denary.Uint64(0),
		},



		{
			Src:      denary.Int8(1),
			Expected: denary.Int8(1),
		},
		{
			Src:      denary.Int16(1),
			Expected: denary.Int16(1),
		},
		{
			Src:      denary.Int32(1),
			Expected: denary.Int32(1),
		},
		{
			Src:      denary.Int64(1),
			Expected: denary.Int64(1),
		},



		{
			Src:      denary.Uint8(1),
			Expected: denary.Uint8(1),
		},
		{
			Src:      denary.Uint16(1),
			Expected: denary.Uint16(1),
		},
		{
			Src:      denary.Uint32(1),
			Expected: denary.Uint32(1),
		},
		{
			Src:      denary.Uint64(1),
			Expected: denary.Uint64(1),
		},



		{
			Src:      denary.Int8(127),
			Expected: denary.Int8(127),
		},
		{
			Src:      denary.Int16(127),
			Expected: denary.Int16(127),
		},
		{
			Src:      denary.Int32(127),
			Expected: denary.Int32(127),
		},
		{
			Src:      denary.Int64(127),
			Expected: denary.Int64(127),
		},



		{
			Src:      denary.Uint8(127),
			Expected: denary.Uint8(127),
		},
		{
			Src:      denary.Uint16(127),
			Expected: denary.Uint16(127),
		},
		{
			Src:      denary.Uint32(127),
			Expected: denary.Uint32(127),
		},
		{
			Src:      denary.Uint64(127),
			Expected: denary.Uint64(127),
		},









		{
			Src:      denary.Int8(-5),
			Expected: denary.Int64(-5),
		},
	}

	for testNumber, test := range tests {

		result := denary.Parse(test.Src)

		actual, ok := result.Unwrap()
		if !ok {
			err := result.Validate()
			t.Errorf("For test #%d, could not unwrap the result (%t), because: (%T) %q", testNumber, ok, err, err)
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
