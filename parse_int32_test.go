package denary_test

import (
	"github.com/reiver/go-denary"

	"testing"
)

func TestParseInt32(t *testing.T) {

	tests := []struct{
		Src int32
		Expected denary.Type
	}{
		{
			Src:             int32(0),
			Expected: denary.Int32(0),
		},
		{
			Src:             int32(1),
			Expected: denary.Int32(1),
		},
		{
			Src:             int32(2),
			Expected: denary.Int32(2),
		},
		{
			Src:             int32(3),
			Expected: denary.Int32(3),
		},
		{
			Src:             int32(4),
			Expected: denary.Int32(4),
		},
		{
			Src:             int32(5),
			Expected: denary.Int32(5),
		},



		{
			Src:             int32(10),
			Expected: denary.Int32(10),
		},



		{
			Src:             int32(20),
			Expected: denary.Int32(20),
		},



		{
			Src:             int32(30),
			Expected: denary.Int32(30),
		},



		{
			Src:             int32(250),
			Expected: denary.Int32(250),
		},
		{
			Src:             int32(251),
			Expected: denary.Int32(251),
		},
		{
			Src:             int32(252),
			Expected: denary.Int32(252),
		},
		{
			Src:             int32(253),
			Expected: denary.Int32(253),
		},
		{
			Src:             int32(254),
			Expected: denary.Int32(254),
		},
		{
			Src:             int32(255),
			Expected: denary.Int32(255),
		},
		{
			Src:             int32(256),
			Expected: denary.Int32(256),
		},
		{
			Src:             int32(257),
			Expected: denary.Int32(257),
		},
		{
			Src:             int32(258),
			Expected: denary.Int32(258),
		},
		{
			Src:             int32(259),
			Expected: denary.Int32(259),
		},
		{
			Src:             int32(261),
			Expected: denary.Int32(261),
		},
		{
			Src:             int32(262),
			Expected: denary.Int32(262),
		},



		{
			Src:             int32(500),
			Expected: denary.Int32(500),
		},



		{
			Src:             int32(1000),
			Expected: denary.Int32(1000),
		},



		{
			Src:             int32(1234),
			Expected: denary.Int32(1234),
		},



		{
			Src:             int32(12345),
			Expected: denary.Int32(12345),
		},



		{
			Src:             int32(65528),
			Expected: denary.Int32(65528),
		},
		{
			Src:             int32(65529),
			Expected: denary.Int32(65529),
		},
		{
			Src:             int32(65530),
			Expected: denary.Int32(65530),
		},
		{
			Src:             int32(65531),
			Expected: denary.Int32(65531),
		},
		{
			Src:             int32(65532),
			Expected: denary.Int32(65532),
		},
		{
			Src:             int32(65533),
			Expected: denary.Int32(65533),
		},
		{
			Src:             int32(65534),
			Expected: denary.Int32(65534),
		},
		{
			Src:             int32(65535),
			Expected: denary.Int32(65535),
		},



		{
			Src:             int32(100000),
			Expected: denary.Int32(100000),
		},



		{
			Src:             int32(1000000),
			Expected: denary.Int32(1000000),
		},



		{
			Src:             int32(1000000),
			Expected: denary.Int32(1000000),
		},



		{
			Src:             int32(10000000),
			Expected: denary.Int32(10000000),
		},



		{
			Src:             int32(100000000),
			Expected: denary.Int32(100000000),
		},




		{
			Src:             int32(1000000000),
			Expected: denary.Int32(1000000000),
		},



		{
			Src:             int32(2000000000),
			Expected: denary.Int32(2000000000),
		},



		{
			Src:             int32(214748364),
			Expected: denary.Int32(214748364),
		},
		{
			Src:             int32(2147483639),
			Expected: denary.Int32(2147483639),
		},
		{
			Src:             int32(2147483640),
			Expected: denary.Int32(2147483640),
		},
		{
			Src:             int32(2147483641),
			Expected: denary.Int32(2147483641),
		},
		{
			Src:             int32(2147483642),
			Expected: denary.Int32(2147483642),
		},
		{
			Src:             int32(2147483643),
			Expected: denary.Int32(2147483643),
		},
		{
			Src:             int32(2147483644),
			Expected: denary.Int32(2147483644),
		},
		{
			Src:             int32(2147483645),
			Expected: denary.Int32(2147483645),
		},
		{
			Src:             int32(2147483646),
			Expected: denary.Int32(2147483646),
		},
		{
			Src:             int32(2147483647),
			Expected: denary.Int32(2147483647),
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
