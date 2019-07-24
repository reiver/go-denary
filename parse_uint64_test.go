package denary_test

import (
	"github.com/reiver/go-denary"

	"testing"
)

func TestParseUint64(t *testing.T) {

	tests := []struct{
		Src uint64
		Expected denary.Type
	}{
		{
			Src:             uint64(0),
			Expected: denary.Uint64(0),
		},
		{
			Src:             uint64(1),
			Expected: denary.Uint64(1),
		},
		{
			Src:             uint64(2),
			Expected: denary.Uint64(2),
		},
		{
			Src:             uint64(3),
			Expected: denary.Uint64(3),
		},
		{
			Src:             uint64(4),
			Expected: denary.Uint64(4),
		},
		{
			Src:             uint64(5),
			Expected: denary.Uint64(5),
		},



		{
			Src:             uint64(10),
			Expected: denary.Uint64(10),
		},



		{
			Src:             uint64(20),
			Expected: denary.Uint64(20),
		},



		{
			Src:             uint64(30),
			Expected: denary.Uint64(30),
		},



		{
			Src:             uint64(250),
			Expected: denary.Uint64(250),
		},
		{
			Src:             uint64(251),
			Expected: denary.Uint64(251),
		},
		{
			Src:             uint64(252),
			Expected: denary.Uint64(252),
		},
		{
			Src:             uint64(253),
			Expected: denary.Uint64(253),
		},
		{
			Src:             uint64(254),
			Expected: denary.Uint64(254),
		},
		{
			Src:             uint64(255),
			Expected: denary.Uint64(255),
		},
		{
			Src:             uint64(256),
			Expected: denary.Uint64(256),
		},
		{
			Src:             uint64(257),
			Expected: denary.Uint64(257),
		},
		{
			Src:             uint64(258),
			Expected: denary.Uint64(258),
		},
		{
			Src:             uint64(259),
			Expected: denary.Uint64(259),
		},
		{
			Src:             uint64(261),
			Expected: denary.Uint64(261),
		},
		{
			Src:             uint64(262),
			Expected: denary.Uint64(262),
		},



		{
			Src:             uint64(500),
			Expected: denary.Uint64(500),
		},



		{
			Src:             uint64(1000),
			Expected: denary.Uint64(1000),
		},



		{
			Src:             uint64(1234),
			Expected: denary.Uint64(1234),
		},



		{
			Src:             uint64(12345),
			Expected: denary.Uint64(12345),
		},



		{
			Src:             uint64(65528),
			Expected: denary.Uint64(65528),
		},
		{
			Src:             uint64(65529),
			Expected: denary.Uint64(65529),
		},
		{
			Src:             uint64(65530),
			Expected: denary.Uint64(65530),
		},
		{
			Src:             uint64(65531),
			Expected: denary.Uint64(65531),
		},
		{
			Src:             uint64(65532),
			Expected: denary.Uint64(65532),
		},
		{
			Src:             uint64(65533),
			Expected: denary.Uint64(65533),
		},
		{
			Src:             uint64(65534),
			Expected: denary.Uint64(65534),
		},
		{
			Src:             uint64(65535),
			Expected: denary.Uint64(65535),
		},



		{
			Src:             uint64(100000),
			Expected: denary.Uint64(100000),
		},



		{
			Src:             uint64(1000000),
			Expected: denary.Uint64(1000000),
		},



		{
			Src:             uint64(1000000),
			Expected: denary.Uint64(1000000),
		},



		{
			Src:             uint64(10000000),
			Expected: denary.Uint64(10000000),
		},



		{
			Src:             uint64(100000000),
			Expected: denary.Uint64(100000000),
		},




		{
			Src:             uint64(1000000000),
			Expected: denary.Uint64(1000000000),
		},



		{
			Src:             uint64(2000000000),
			Expected: denary.Uint64(2000000000),
		},



		{
			Src:             uint64(3000000000),
			Expected: denary.Uint64(3000000000),
		},



		{
			Src:             uint64(4000000000),
			Expected: denary.Uint64(4000000000),
		},



		{
			Src:             uint64(4294967288),
			Expected: denary.Uint64(4294967288),
		},
		{
			Src:             uint64(4294967289),
			Expected: denary.Uint64(4294967289),
		},
		{
			Src:             uint64(4294967290),
			Expected: denary.Uint64(4294967290),
		},
		{
			Src:             uint64(4294967291),
			Expected: denary.Uint64(4294967291),
		},
		{
			Src:             uint64(4294967292),
			Expected: denary.Uint64(4294967292),
		},
		{
			Src:             uint64(4294967293),
			Expected: denary.Uint64(4294967293),
		},
		{
			Src:             uint64(4294967294),
			Expected: denary.Uint64(4294967294),
		},
		{
			Src:             uint64(4294967295),
			Expected: denary.Uint64(4294967295),
		},



		{
			Src:             uint64(5000000000),
			Expected: denary.Uint64(5000000000),
		},



		{
			Src:             uint64(10000000000),
			Expected: denary.Uint64(10000000000),
		},



		{
			Src:             uint64(100000000000),
			Expected: denary.Uint64(100000000000),
		},



		{
			Src:             uint64(1000000000000),
			Expected: denary.Uint64(1000000000000),
		},



		{
			Src:             uint64(10000000000000),
			Expected: denary.Uint64(10000000000000),
		},



		{
			Src:             uint64(100000000000000),
			Expected: denary.Uint64(100000000000000),
		},



		{
			Src:             uint64(1000000000000000),
			Expected: denary.Uint64(1000000000000000),
		},



		{
			Src:             uint64(10000000000000000),
			Expected: denary.Uint64(10000000000000000),
		},



		{
			Src:             uint64(100000000000000000),
			Expected: denary.Uint64(100000000000000000),
		},



		{
			Src:             uint64(1000000000000000000),
			Expected: denary.Uint64(1000000000000000000),
		},



		{
			Src:             uint64(9223372036854775807),
			Expected: denary.Uint64(9223372036854775807),
		},



		{
			Src:             uint64(10000000000000000000),
			Expected: denary.Uint64(10000000000000000000),
		},



		{
			Src:             uint64(18446744073709551615),
			Expected: denary.Uint64(18446744073709551615),
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
