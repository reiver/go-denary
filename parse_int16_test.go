package denary_test

import (
	"github.com/reiver/go-denary"

	"testing"
)

func TestParseInt16(t *testing.T) {

	tests := []struct{
		Src int16
		Expected denary.Type
	}{
		{
			Src:             int16(0),
			Expected: denary.Int16(0),
		},
		{
			Src:             int16(1),
			Expected: denary.Int16(1),
		},
		{
			Src:             int16(2),
			Expected: denary.Int16(2),
		},
		{
			Src:             int16(3),
			Expected: denary.Int16(3),
		},
		{
			Src:             int16(4),
			Expected: denary.Int16(4),
		},
		{
			Src:             int16(5),
			Expected: denary.Int16(5),
		},



		{
			Src:             int16(10),
			Expected: denary.Int16(10),
		},



		{
			Src:             int16(20),
			Expected: denary.Int16(20),
		},



		{
			Src:             int16(30),
			Expected: denary.Int16(30),
		},



		{
			Src:             int16(250),
			Expected: denary.Int16(250),
		},
		{
			Src:             int16(251),
			Expected: denary.Int16(251),
		},
		{
			Src:             int16(252),
			Expected: denary.Int16(252),
		},
		{
			Src:             int16(253),
			Expected: denary.Int16(253),
		},
		{
			Src:             int16(254),
			Expected: denary.Int16(254),
		},
		{
			Src:             int16(255),
			Expected: denary.Int16(255),
		},
		{
			Src:             int16(256),
			Expected: denary.Int16(256),
		},
		{
			Src:             int16(257),
			Expected: denary.Int16(257),
		},
		{
			Src:             int16(258),
			Expected: denary.Int16(258),
		},
		{
			Src:             int16(259),
			Expected: denary.Int16(259),
		},
		{
			Src:             int16(261),
			Expected: denary.Int16(261),
		},
		{
			Src:             int16(262),
			Expected: denary.Int16(262),
		},



		{
			Src:             int16(500),
			Expected: denary.Int16(500),
		},



		{
			Src:             int16(1000),
			Expected: denary.Int16(1000),
		},



		{
			Src:             int16(1234),
			Expected: denary.Int16(1234),
		},



		{
			Src:             int16(12345),
			Expected: denary.Int16(12345),
		},



		{
			Src:             int16(32759),
			Expected: denary.Int16(32759),
		},
		{
			Src:             int16(32760),
			Expected: denary.Int16(32760),
		},
		{
			Src:             int16(32761),
			Expected: denary.Int16(32761),
		},
		{
			Src:             int16(32762),
			Expected: denary.Int16(32762),
		},
		{
			Src:             int16(32763),
			Expected: denary.Int16(32763),
		},
		{
			Src:             int16(32764),
			Expected: denary.Int16(32764),
		},
		{
			Src:             int16(32765),
			Expected: denary.Int16(32765),
		},
		{
			Src:             int16(32766),
			Expected: denary.Int16(32766),
		},
		{
			Src:             int16(32767),
			Expected: denary.Int16(32767),
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
