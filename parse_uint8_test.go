package denary_test

import (
	"github.com/reiver/go-denary"

	"testing"
)

func TestParseUint8(t *testing.T) {

	tests := []struct{
		Src uint8
		Expected denary.Type
	}{
		{
			Src:             uint8(0),
			Expected: denary.Uint8(0),
		},
		{
			Src:             uint8(1),
			Expected: denary.Uint8(1),
		},
		{
			Src:             uint8(2),
			Expected: denary.Uint8(2),
		},
		{
			Src:             uint8(3),
			Expected: denary.Uint8(3),
		},
		{
			Src:             uint8(4),
			Expected: denary.Uint8(4),
		},
		{
			Src:             uint8(5),
			Expected: denary.Uint8(5),
		},



		{
			Src:             uint8(10),
			Expected: denary.Uint8(10),
		},



		{
			Src:             uint8(20),
			Expected: denary.Uint8(20),
		},



		{
			Src:             uint8(30),
			Expected: denary.Uint8(30),
		},



		{
			Src:             uint8(250),
			Expected: denary.Uint8(250),
		},
		{
			Src:             uint8(251),
			Expected: denary.Uint8(251),
		},
		{
			Src:             uint8(252),
			Expected: denary.Uint8(252),
		},
		{
			Src:             uint8(253),
			Expected: denary.Uint8(253),
		},
		{
			Src:             uint8(254),
			Expected: denary.Uint8(254),
		},
		{
			Src:             uint8(255),
			Expected: denary.Uint8(255),
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
