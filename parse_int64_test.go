package denary_test

import (
	"github.com/reiver/go-denary"

	"testing"
)

func TestParseInt64(t *testing.T) {

	tests := []struct{
		Src int64
		Expected denary.Type
	}{
		{
			Src:             int64(0),
			Expected: denary.Int64(0),
		},
		{
			Src:             int64(1),
			Expected: denary.Int64(1),
		},
		{
			Src:             int64(2),
			Expected: denary.Int64(2),
		},
		{
			Src:             int64(3),
			Expected: denary.Int64(3),
		},
		{
			Src:             int64(4),
			Expected: denary.Int64(4),
		},
		{
			Src:             int64(5),
			Expected: denary.Int64(5),
		},



		{
			Src:             int64(10),
			Expected: denary.Int64(10),
		},



		{
			Src:             int64(20),
			Expected: denary.Int64(20),
		},



		{
			Src:             int64(30),
			Expected: denary.Int64(30),
		},



		{
			Src:             int64(250),
			Expected: denary.Int64(250),
		},
		{
			Src:             int64(251),
			Expected: denary.Int64(251),
		},
		{
			Src:             int64(252),
			Expected: denary.Int64(252),
		},
		{
			Src:             int64(253),
			Expected: denary.Int64(253),
		},
		{
			Src:             int64(254),
			Expected: denary.Int64(254),
		},
		{
			Src:             int64(255),
			Expected: denary.Int64(255),
		},
		{
			Src:             int64(256),
			Expected: denary.Int64(256),
		},
		{
			Src:             int64(257),
			Expected: denary.Int64(257),
		},
		{
			Src:             int64(258),
			Expected: denary.Int64(258),
		},
		{
			Src:             int64(259),
			Expected: denary.Int64(259),
		},
		{
			Src:             int64(261),
			Expected: denary.Int64(261),
		},
		{
			Src:             int64(262),
			Expected: denary.Int64(262),
		},



		{
			Src:             int64(500),
			Expected: denary.Int64(500),
		},



		{
			Src:             int64(1000),
			Expected: denary.Int64(1000),
		},



		{
			Src:             int64(1234),
			Expected: denary.Int64(1234),
		},



		{
			Src:             int64(12345),
			Expected: denary.Int64(12345),
		},



		{
			Src:             int64(65528),
			Expected: denary.Int64(65528),
		},
		{
			Src:             int64(65529),
			Expected: denary.Int64(65529),
		},
		{
			Src:             int64(65530),
			Expected: denary.Int64(65530),
		},
		{
			Src:             int64(65531),
			Expected: denary.Int64(65531),
		},
		{
			Src:             int64(65532),
			Expected: denary.Int64(65532),
		},
		{
			Src:             int64(65533),
			Expected: denary.Int64(65533),
		},
		{
			Src:             int64(65534),
			Expected: denary.Int64(65534),
		},
		{
			Src:             int64(65535),
			Expected: denary.Int64(65535),
		},



		{
			Src:             int64(100000),
			Expected: denary.Int64(100000),
		},



		{
			Src:             int64(1000000),
			Expected: denary.Int64(1000000),
		},



		{
			Src:             int64(1000000),
			Expected: denary.Int64(1000000),
		},



		{
			Src:             int64(10000000),
			Expected: denary.Int64(10000000),
		},



		{
			Src:             int64(100000000),
			Expected: denary.Int64(100000000),
		},




		{
			Src:             int64(1000000000),
			Expected: denary.Int64(1000000000),
		},



		{
			Src:             int64(2000000000),
			Expected: denary.Int64(2000000000),
		},



		{
			Src:             int64(3000000000),
			Expected: denary.Int64(3000000000),
		},



		{
			Src:             int64(4000000000),
			Expected: denary.Int64(4000000000),
		},



		{
			Src:             int64(4294967288),
			Expected: denary.Int64(4294967288),
		},
		{
			Src:             int64(4294967289),
			Expected: denary.Int64(4294967289),
		},
		{
			Src:             int64(4294967290),
			Expected: denary.Int64(4294967290),
		},
		{
			Src:             int64(4294967291),
			Expected: denary.Int64(4294967291),
		},
		{
			Src:             int64(4294967292),
			Expected: denary.Int64(4294967292),
		},
		{
			Src:             int64(4294967293),
			Expected: denary.Int64(4294967293),
		},
		{
			Src:             int64(4294967294),
			Expected: denary.Int64(4294967294),
		},
		{
			Src:             int64(4294967295),
			Expected: denary.Int64(4294967295),
		},



		{
			Src:             int64(5000000000),
			Expected: denary.Int64(5000000000),
		},



		{
			Src:             int64(10000000000),
			Expected: denary.Int64(10000000000),
		},



		{
			Src:             int64(100000000000),
			Expected: denary.Int64(100000000000),
		},



		{
			Src:             int64(1000000000000),
			Expected: denary.Int64(1000000000000),
		},



		{
			Src:             int64(10000000000000),
			Expected: denary.Int64(10000000000000),
		},



		{
			Src:             int64(100000000000000),
			Expected: denary.Int64(100000000000000),
		},



		{
			Src:             int64(1000000000000000),
			Expected: denary.Int64(1000000000000000),
		},



		{
			Src:             int64(10000000000000000),
			Expected: denary.Int64(10000000000000000),
		},



		{
			Src:             int64(100000000000000000),
			Expected: denary.Int64(100000000000000000),
		},



		{
			Src:             int64(1000000000000000000),
			Expected: denary.Int64(1000000000000000000),
		},



		{
			Src:             int64(9223372036854775779),
			Expected: denary.Int64(9223372036854775779),
		},
		{
			Src:             int64(9223372036854775800),
			Expected: denary.Int64(9223372036854775800),
		},
		{
			Src:             int64(9223372036854775801),
			Expected: denary.Int64(9223372036854775801),
		},
		{
			Src:             int64(9223372036854775802),
			Expected: denary.Int64(9223372036854775802),
		},
		{
			Src:             int64(9223372036854775803),
			Expected: denary.Int64(9223372036854775803),
		},
		{
			Src:             int64(9223372036854775804),
			Expected: denary.Int64(9223372036854775804),
		},
		{
			Src:             int64(9223372036854775805),
			Expected: denary.Int64(9223372036854775805),
		},
		{
			Src:             int64(9223372036854775806),
			Expected: denary.Int64(9223372036854775806),
		},
		{
			Src:             int64(9223372036854775807),
			Expected: denary.Int64(9223372036854775807),
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
