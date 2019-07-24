package denary_test

import (
	"github.com/reiver/go-denary"

	"testing"
)

func TestParseInt8(t *testing.T) {

	tests := []struct{
		Src int8
		Expected denary.Type
	}{
		{
			Src:             int8(0),
			Expected: denary.Int8(0),
		},
		{
			Src:             int8(1),
			Expected: denary.Int8(1),
		},
		{
			Src:             int8(2),
			Expected: denary.Int8(2),
		},
		{
			Src:             int8(3),
			Expected: denary.Int8(3),
		},
		{
			Src:             int8(4),
			Expected: denary.Int8(4),
		},
		{
			Src:             int8(5),
			Expected: denary.Int8(5),
		},



		{
			Src:             int8(10),
			Expected: denary.Int8(10),
		},



		{
			Src:             int8(20),
			Expected: denary.Int8(20),
		},



		{
			Src:             int8(30),
			Expected: denary.Int8(30),
		},



		{
			Src:             int8(119),
			Expected: denary.Int8(119),
		},
		{
			Src:             int8(120),
			Expected: denary.Int8(120),
		},
		{
			Src:             int8(121),
			Expected: denary.Int8(121),
		},
		{
			Src:             int8(122),
			Expected: denary.Int8(122),
		},
		{
			Src:             int8(123),
			Expected: denary.Int8(123),
		},
		{
			Src:             int8(124),
			Expected: denary.Int8(124),
		},
		{
			Src:             int8(125),
			Expected: denary.Int8(125),
		},
		{
			Src:             int8(126),
			Expected: denary.Int8(126),
		},
		{
			Src:             int8(126),
			Expected: denary.Int8(126),
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
