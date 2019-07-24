package denary_test

import (
	"github.com/reiver/go-denary"

	"testing"
)

func TestParseUint16(t *testing.T) {

	tests := []struct{
		Src uint16
		Expected denary.Type
	}{
		{
			Src:             uint16(0),
			Expected: denary.Uint16(0),
		},
		{
			Src:             uint16(1),
			Expected: denary.Uint16(1),
		},
		{
			Src:             uint16(2),
			Expected: denary.Uint16(2),
		},
		{
			Src:             uint16(3),
			Expected: denary.Uint16(3),
		},
		{
			Src:             uint16(4),
			Expected: denary.Uint16(4),
		},
		{
			Src:             uint16(5),
			Expected: denary.Uint16(5),
		},



		{
			Src:             uint16(10),
			Expected: denary.Uint16(10),
		},



		{
			Src:             uint16(20),
			Expected: denary.Uint16(20),
		},



		{
			Src:             uint16(30),
			Expected: denary.Uint16(30),
		},



		{
			Src:             uint16(250),
			Expected: denary.Uint16(250),
		},
		{
			Src:             uint16(251),
			Expected: denary.Uint16(251),
		},
		{
			Src:             uint16(252),
			Expected: denary.Uint16(252),
		},
		{
			Src:             uint16(253),
			Expected: denary.Uint16(253),
		},
		{
			Src:             uint16(254),
			Expected: denary.Uint16(254),
		},
		{
			Src:             uint16(255),
			Expected: denary.Uint16(255),
		},
		{
			Src:             uint16(256),
			Expected: denary.Uint16(256),
		},
		{
			Src:             uint16(257),
			Expected: denary.Uint16(257),
		},
		{
			Src:             uint16(258),
			Expected: denary.Uint16(258),
		},
		{
			Src:             uint16(259),
			Expected: denary.Uint16(259),
		},
		{
			Src:             uint16(261),
			Expected: denary.Uint16(261),
		},
		{
			Src:             uint16(262),
			Expected: denary.Uint16(262),
		},



		{
			Src:             uint16(500),
			Expected: denary.Uint16(500),
		},



		{
			Src:             uint16(1000),
			Expected: denary.Uint16(1000),
		},



		{
			Src:             uint16(1234),
			Expected: denary.Uint16(1234),
		},



		{
			Src:             uint16(12345),
			Expected: denary.Uint16(12345),
		},



		{
			Src:             uint16(65528),
			Expected: denary.Uint16(65528),
		},
		{
			Src:             uint16(65529),
			Expected: denary.Uint16(65529),
		},
		{
			Src:             uint16(65530),
			Expected: denary.Uint16(65530),
		},
		{
			Src:             uint16(65531),
			Expected: denary.Uint16(65531),
		},
		{
			Src:             uint16(65532),
			Expected: denary.Uint16(65532),
		},
		{
			Src:             uint16(65533),
			Expected: denary.Uint16(65533),
		},
		{
			Src:             uint16(65534),
			Expected: denary.Uint16(65534),
		},
		{
			Src:             uint16(65535),
			Expected: denary.Uint16(65535),
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
