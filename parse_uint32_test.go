package denary_test

import (
	"github.com/reiver/go-denary"

	"testing"
)

func TestParseUint32(t *testing.T) {

	tests := []struct{
		Src uint32
		Expected denary.Type
	}{
		{
			Src:             uint32(0),
			Expected: denary.Uint32(0),
		},
		{
			Src:             uint32(1),
			Expected: denary.Uint32(1),
		},
		{
			Src:             uint32(2),
			Expected: denary.Uint32(2),
		},
		{
			Src:             uint32(3),
			Expected: denary.Uint32(3),
		},
		{
			Src:             uint32(4),
			Expected: denary.Uint32(4),
		},
		{
			Src:             uint32(5),
			Expected: denary.Uint32(5),
		},



		{
			Src:             uint32(10),
			Expected: denary.Uint32(10),
		},



		{
			Src:             uint32(20),
			Expected: denary.Uint32(20),
		},



		{
			Src:             uint32(30),
			Expected: denary.Uint32(30),
		},



		{
			Src:             uint32(250),
			Expected: denary.Uint32(250),
		},
		{
			Src:             uint32(251),
			Expected: denary.Uint32(251),
		},
		{
			Src:             uint32(252),
			Expected: denary.Uint32(252),
		},
		{
			Src:             uint32(253),
			Expected: denary.Uint32(253),
		},
		{
			Src:             uint32(254),
			Expected: denary.Uint32(254),
		},
		{
			Src:             uint32(255),
			Expected: denary.Uint32(255),
		},
		{
			Src:             uint32(256),
			Expected: denary.Uint32(256),
		},
		{
			Src:             uint32(257),
			Expected: denary.Uint32(257),
		},
		{
			Src:             uint32(258),
			Expected: denary.Uint32(258),
		},
		{
			Src:             uint32(259),
			Expected: denary.Uint32(259),
		},
		{
			Src:             uint32(261),
			Expected: denary.Uint32(261),
		},
		{
			Src:             uint32(262),
			Expected: denary.Uint32(262),
		},



		{
			Src:             uint32(500),
			Expected: denary.Uint32(500),
		},



		{
			Src:             uint32(1000),
			Expected: denary.Uint32(1000),
		},



		{
			Src:             uint32(1234),
			Expected: denary.Uint32(1234),
		},



		{
			Src:             uint32(12345),
			Expected: denary.Uint32(12345),
		},



		{
			Src:             uint32(65528),
			Expected: denary.Uint32(65528),
		},
		{
			Src:             uint32(65529),
			Expected: denary.Uint32(65529),
		},
		{
			Src:             uint32(65530),
			Expected: denary.Uint32(65530),
		},
		{
			Src:             uint32(65531),
			Expected: denary.Uint32(65531),
		},
		{
			Src:             uint32(65532),
			Expected: denary.Uint32(65532),
		},
		{
			Src:             uint32(65533),
			Expected: denary.Uint32(65533),
		},
		{
			Src:             uint32(65534),
			Expected: denary.Uint32(65534),
		},
		{
			Src:             uint32(65535),
			Expected: denary.Uint32(65535),
		},



		{
			Src:             uint32(100000),
			Expected: denary.Uint32(100000),
		},



		{
			Src:             uint32(1000000),
			Expected: denary.Uint32(1000000),
		},



		{
			Src:             uint32(1000000),
			Expected: denary.Uint32(1000000),
		},



		{
			Src:             uint32(10000000),
			Expected: denary.Uint32(10000000),
		},



		{
			Src:             uint32(100000000),
			Expected: denary.Uint32(100000000),
		},




		{
			Src:             uint32(1000000000),
			Expected: denary.Uint32(1000000000),
		},



		{
			Src:             uint32(2000000000),
			Expected: denary.Uint32(2000000000),
		},



		{
			Src:             uint32(3000000000),
			Expected: denary.Uint32(3000000000),
		},



		{
			Src:             uint32(4000000000),
			Expected: denary.Uint32(4000000000),
		},



		{
			Src:             uint32(4294967288),
			Expected: denary.Uint32(4294967288),
		},
		{
			Src:             uint32(4294967289),
			Expected: denary.Uint32(4294967289),
		},
		{
			Src:             uint32(4294967290),
			Expected: denary.Uint32(4294967290),
		},
		{
			Src:             uint32(4294967291),
			Expected: denary.Uint32(4294967291),
		},
		{
			Src:             uint32(4294967292),
			Expected: denary.Uint32(4294967292),
		},
		{
			Src:             uint32(4294967293),
			Expected: denary.Uint32(4294967293),
		},
		{
			Src:             uint32(4294967294),
			Expected: denary.Uint32(4294967294),
		},
		{
			Src:             uint32(4294967295),
			Expected: denary.Uint32(4294967295),
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
