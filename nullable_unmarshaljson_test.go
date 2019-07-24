package denary_test

import (
	"github.com/reiver/go-denary"

	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"testing"
)

func TestNullableUnmarshalJSONNull(t *testing.T) {

	{
		var actualNullable denary.Nullable

		if expected, actual := denary.Nothing().Nullable(), actualNullable; expected != actual {
			t.Errorf("The denary.Nullable actually had the value %#v, not the expected value of %#v", actual, expected)
			return
		}

		var jsonString string = "null"

		if err := json.Unmarshal([]byte(jsonString), &actualNullable); nil != err {
			t.Errorf("Did not expect an error, for calling json.Unmarshal(), but actually got one: (%T) %q", err, err)
			return
		}
		if expected, actual := denary.Null(), actualNullable; expected != actual {
			t.Errorf("The denary.Nullable actually had the value %#v, not the expected value of %#v", actual, expected)
			return
		}
	}

	{
		type MyStruct struct {
			This string          `json:"this"`
			That denary.Nullable `json:"that"`
		}

		var actualStruct MyStruct

		if expected, actual := denary.Nothing().Nullable(), actualStruct.That; expected != actual {
			t.Errorf("The denary.Nullable actually had the value %#v, not the expected value of %#v", actual, expected)
			return
		}

		var jsonString string =
`{
	"this":"Hello world!",
	"that":null
}`

		if err := json.Unmarshal([]byte(jsonString), &actualStruct); nil != err {
			t.Errorf("Did not expect an error, for calling json.Unmarshal(), but actually got one: (%T) %q", err, err)
			return
		}
		if expected, actual := denary.Null(), actualStruct.That; expected != actual {
			t.Errorf("The denary.Nullable actually had the value %#v, not the expected value of %#v", actual, expected)
			return
		}
	}
}

func TestNullableUnmarshalJSON(t *testing.T) {

	tests := []struct{
		Src      string
		Expected string
	}{
		{
			Src:      "-.8904",
			Expected: "-.8904",
		},









		{
			Src:      "-127.",
			Expected: "-127.",
		},

		{
			Src:      "-5.",
			Expected: "-5.",
		},
		{
			Src:      "-4.",
			Expected: "-4.",
		},
		{
			Src:      "-3.",
			Expected: "-3.",
		},
		{
			Src:      "-2.",
			Expected: "-2.",
		},
		{
			Src:      "-1.",
			Expected: "-1.",
		},
		{
			Src:      "-0.",
			Expected: "-0.",
		},
		{
			Src:      "0.",
			Expected: "0.",
		},
		{
			Src:      "1.",
			Expected: "1.",
		},
		{
			Src:      "2.",
			Expected: "2.",
		},
		{
			Src:      "3.",
			Expected: "3.",
		},
		{
			Src:      "4.",
			Expected: "4.",
		},
		{
			Src:      "5.",
			Expected: "5.",
		},

		{
			Src:      "127.",
			Expected: "127.",
		},









		{
			Src:      "-.127",
			Expected: "-.127",
		},

		{
			Src:      "-.5",
			Expected: "-.5",
		},
		{
			Src:      "-.4",
			Expected: "-.4",
		},
		{
			Src:      "-.3",
			Expected: "-.3",
		},
		{
			Src:      "-.2",
			Expected: "-.2",
		},
		{
			Src:      "-.1",
			Expected: "-.1",
		},
		{
			Src:      "-.0",
			Expected: "-.0",
		},
		{
			Src:      ".0",
			Expected: ".0",
		},
		{
			Src:      ".1",
			Expected: ".1",
		},
		{
			Src:      ".2",
			Expected: ".2",
		},
		{
			Src:      ".3",
			Expected: ".3",
		},
		{
			Src:      ".4",
			Expected: ".4",
		},
		{
			Src:      ".5",
			Expected: ".5",
		},









		{
			Src:      "-39614081257132168796771975168.604462909807314587353088",
			Expected: "-39614081257132168796771975168.604462909807314587353088",
		},

		{
			Src:      "-5.5",
			Expected: "-5.5",
		},
		{
			Src:      "-4.4",
			Expected: "-4.4",
		},
		{
			Src:      "-3.3",
			Expected: "-3.3",
		},
		{
			Src:      "-2.2",
			Expected: "-2.2",
		},
		{
			Src:      "-1.1",
			Expected: "-1.1",
		},
		{
			Src:      "-0.0",
			Expected: "-0.0",
		},
		{
			Src:      "0.0",
			Expected: "0.0",
		},
		{
			Src:      "1.1",
			Expected: "1.1",
		},
		{
			Src:      "2.2",
			Expected: "2.2",
		},
		{
			Src:      "3.3",
			Expected: "3.3",
		},
		{
			Src:      "4.4",
			Expected: "4.4",
		},
		{
			Src:      "5.5",
			Expected: "5.5",
		},

		{
			Src:      "39614081257132168796771975168.604462909807314587353088",
			Expected: "39614081257132168796771975168.604462909807314587353088",
		},









		{
			Src:      "-5",
			Expected: "-5",
		},



//@TODO
//		{
//			Src:      "-1,234,567.890123456",
//			Expected: "-1234567.890123456",
//		},



		{
			Src:      "-128",
			Expected: "-128",
		},



		{
			Src:      "-3.14159265358979323846264338327950",
			Expected: "-3.14159265358979323846264338327950",
		},



		{
			Src:      "-1",
			Expected: "-1",
		},



		{
			Src:      "-0.12345678901",
			Expected: "-0.12345678901",
		},



		{
			Src:      "-1.0",
			Expected: "-1.0",
		},
		{
			Src:      "-0.9",
			Expected: "-0.9",
		},
		{
			Src:      "-0.8",
			Expected: "-0.8",
		},
		{
			Src:      "-0.7",
			Expected: "-0.7",
		},
		{
			Src:      "-0.6",
			Expected: "-0.6",
		},
		{
			Src:      "-0.5",
			Expected: "-0.5",
		},
		{
			Src:      "-0.4",
			Expected: "-0.4",
		},
		{
			Src:      "-0.3",
			Expected: "-0.3",
		},
		{
			Src:      "-0.2",
			Expected: "-0.2",
		},
		{
			Src:      "-0.1",
			Expected: "-0.1",
		},
		{
			Src:      "-0.0",
			Expected: "-0.0",
		},



		{
			Src:      "-1.00",
			Expected: "-1.00",
		},
		{
			Src:      "-0.99",
			Expected: "-0.99",
		},
		{
			Src:      "-0.98",
			Expected: "-0.98",
		},
		{
			Src:      "-0.97",
			Expected: "-0.97",
		},
		{
			Src:      "-0.96",
			Expected: "-0.96",
		},
		{
			Src:      "-0.95",
			Expected: "-0.95",
		},
		{
			Src:      "-0.94",
			Expected: "-0.94",
		},
		{
			Src:      "-0.93",
			Expected: "-0.93",
		},
		{
			Src:      "-0.92",
			Expected: "-0.92",
		},
		{
			Src:      "-0.91",
			Expected: "-0.91",
		},
		{
			Src:      "-0.90",
			Expected: "-0.90",
		},
		{
			Src:      "-0.89",
			Expected: "-0.89",
		},
		{
			Src:      "-0.88",
			Expected: "-0.88",
		},
		{
			Src:      "-0.87",
			Expected: "-0.87",
		},
		{
			Src:      "-0.86",
			Expected: "-0.86",
		},
		{
			Src:      "-0.85",
			Expected: "-0.85",
		},
		{
			Src:      "-0.84",
			Expected: "-0.84",
		},
		{
			Src:      "-0.83",
			Expected: "-0.83",
		},
		{
			Src:      "-0.82",
			Expected: "-0.82",
		},
		{
			Src:      "-0.81",
			Expected: "-0.81",
		},
		{
			Src:      "-0.80",
			Expected: "-0.80",
		},
		{
			Src:      "-0.79",
			Expected: "-0.79",
		},
		{
			Src:      "-0.78",
			Expected: "-0.78",
		},
		{
			Src:      "-0.77",
			Expected: "-0.77",
		},
		{
			Src:      "-0.76",
			Expected: "-0.76",
		},
		{
			Src:      "-0.75",
			Expected: "-0.75",
		},
		{
			Src:      "-0.74",
			Expected: "-0.74",
		},
		{
			Src:      "-0.73",
			Expected: "-0.73",
		},
		{
			Src:      "-0.72",
			Expected: "-0.72",
		},
		{
			Src:      "-0.71",
			Expected: "-0.71",
		},
		{
			Src:      "-0.70",
			Expected: "-0.70",
		},
		{
			Src:      "-0.69",
			Expected: "-0.69",
		},
		{
			Src:      "-0.68",
			Expected: "-0.68",
		},
		{
			Src:      "-0.67",
			Expected: "-0.67",
		},
		{
			Src:      "-0.66",
			Expected: "-0.66",
		},
		{
			Src:      "-0.65",
			Expected: "-0.65",
		},
		{
			Src:      "-0.64",
			Expected: "-0.64",
		},
		{
			Src:      "-0.63",
			Expected: "-0.63",
		},
		{
			Src:      "-0.62",
			Expected: "-0.62",
		},
		{
			Src:      "-0.61",
			Expected: "-0.61",
		},
		{
			Src:      "-0.60",
			Expected: "-0.60",
		},
		{
			Src:      "-0.59",
			Expected: "-0.59",
		},
		{
			Src:      "-0.58",
			Expected: "-0.58",
		},
		{
			Src:      "-0.57",
			Expected: "-0.57",
		},
		{
			Src:      "-0.56",
			Expected: "-0.56",
		},
		{
			Src:      "-0.55",
			Expected: "-0.55",
		},
		{
			Src:      "-0.54",
			Expected: "-0.54",
		},
		{
			Src:      "-0.53",
			Expected: "-0.53",
		},
		{
			Src:      "-0.52",
			Expected: "-0.52",
		},
		{
			Src:      "-0.51",
			Expected: "-0.51",
		},
		{
			Src:      "-0.50",
			Expected: "-0.50",
		},
		{
			Src:      "-0.49",
			Expected: "-0.49",
		},
		{
			Src:      "-0.48",
			Expected: "-0.48",
		},
		{
			Src:      "-0.47",
			Expected: "-0.47",
		},
		{
			Src:      "-0.46",
			Expected: "-0.46",
		},
		{
			Src:      "-0.45",
			Expected: "-0.45",
		},
		{
			Src:      "-0.44",
			Expected: "-0.44",
		},
		{
			Src:      "-0.43",
			Expected: "-0.43",
		},
		{
			Src:      "-0.42",
			Expected: "-0.42",
		},
		{
			Src:      "-0.41",
			Expected: "-0.41",
		},
		{
			Src:      "-0.40",
			Expected: "-0.40",
		},
		{
			Src:      "-0.39",
			Expected: "-0.39",
		},
		{
			Src:      "-0.38",
			Expected: "-0.38",
		},
		{
			Src:      "-0.37",
			Expected: "-0.37",
		},
		{
			Src:      "-0.36",
			Expected: "-0.36",
		},
		{
			Src:      "-0.35",
			Expected: "-0.35",
		},
		{
			Src:      "-0.34",
			Expected: "-0.34",
		},
		{
			Src:      "-0.33",
			Expected: "-0.33",
		},
		{
			Src:      "-0.32",
			Expected: "-0.32",
		},
		{
			Src:      "-0.31",
			Expected: "-0.31",
		},
		{
			Src:      "-0.30",
			Expected: "-0.30",
		},
		{
			Src:      "-0.29",
			Expected: "-0.29",
		},
		{
			Src:      "-0.28",
			Expected: "-0.28",
		},
		{
			Src:      "-0.27",
			Expected: "-0.27",
		},
		{
			Src:      "-0.26",
			Expected: "-0.26",
		},
		{
			Src:      "-0.25",
			Expected: "-0.25",
		},
		{
			Src:      "-0.24",
			Expected: "-0.24",
		},
		{
			Src:      "-0.23",
			Expected: "-0.23",
		},
		{
			Src:      "-0.22",
			Expected: "-0.22",
		},
		{
			Src:      "-0.21",
			Expected: "-0.21",
		},
		{
			Src:      "-0.20",
			Expected: "-0.20",
		},
		{
			Src:      "-0.19",
			Expected: "-0.19",
		},
		{
			Src:      "-0.18",
			Expected: "-0.18",
		},
		{
			Src:      "-0.17",
			Expected: "-0.17",
		},
		{
			Src:      "-0.16",
			Expected: "-0.16",
		},
		{
			Src:      "-0.15",
			Expected: "-0.15",
		},
		{
			Src:      "-0.14",
			Expected: "-0.14",
		},
		{
			Src:      "-0.13",
			Expected: "-0.13",
		},
		{
			Src:      "-0.12",
			Expected: "-0.12",
		},
		{
			Src:      "-0.11",
			Expected: "-0.11",
		},
		{
			Src:      "-0.10",
			Expected: "-0.10",
		},
		{
			Src:      "-0.09",
			Expected: "-0.09",
		},
		{
			Src:      "-0.08",
			Expected: "-0.08",
		},
		{
			Src:      "-0.07",
			Expected: "-0.07",
		},
		{
			Src:      "-0.06",
			Expected: "-0.06",
		},
		{
			Src:      "-0.05",
			Expected: "-0.05",
		},
		{
			Src:      "-0.04",
			Expected: "-0.04",
		},
		{
			Src:      "-0.03",
			Expected: "-0.03",
		},
		{
			Src:      "-0.02",
			Expected: "-0.02",
		},
		{
			Src:      "-0.01",
			Expected: "-0.01",
		},
		{
			Src:      "-0.00",
			Expected: "-0.00",
		},



		{
			Src:      "-0",
			Expected: "-0",
		},
		{
			Src:      "0",
			Expected: "0",
		},



		{
			Src:      "0.00",
			Expected: "0.00",
		},
		{
			Src:      "0.01",
			Expected: "0.01",
		},
		{
			Src:      "0.02",
			Expected: "0.02",
		},
		{
			Src:      "0.03",
			Expected: "0.03",
		},
		{
			Src:      "0.04",
			Expected: "0.04",
		},
		{
			Src:      "0.05",
			Expected: "0.05",
		},
		{
			Src:      "0.06",
			Expected: "0.06",
		},
		{
			Src:      "0.07",
			Expected: "0.07",
		},
		{
			Src:      "0.08",
			Expected: "0.08",
		},
		{
			Src:      "0.09",
			Expected: "0.09",
		},
		{
			Src:      "0.10",
			Expected: "0.10",
		},
		{
			Src:      "0.11",
			Expected: "0.11",
		},
		{
			Src:      "0.12",
			Expected: "0.12",
		},
		{
			Src:      "0.13",
			Expected: "0.13",
		},
		{
			Src:      "0.14",
			Expected: "0.14",
		},
		{
			Src:      "0.15",
			Expected: "0.15",
		},
		{
			Src:      "0.16",
			Expected: "0.16",
		},
		{
			Src:      "0.17",
			Expected: "0.17",
		},
		{
			Src:      "0.18",
			Expected: "0.18",
		},
		{
			Src:      "0.19",
			Expected: "0.19",
		},
		{
			Src:      "0.20",
			Expected: "0.20",
		},
		{
			Src:      "0.21",
			Expected: "0.21",
		},
		{
			Src:      "0.22",
			Expected: "0.22",
		},
		{
			Src:      "0.23",
			Expected: "0.23",
		},
		{
			Src:      "0.24",
			Expected: "0.24",
		},
		{
			Src:      "0.25",
			Expected: "0.25",
		},
		{
			Src:      "0.26",
			Expected: "0.26",
		},
		{
			Src:      "0.27",
			Expected: "0.27",
		},
		{
			Src:      "0.28",
			Expected: "0.28",
		},
		{
			Src:      "0.29",
			Expected: "0.29",
		},
		{
			Src:      "0.30",
			Expected: "0.30",
		},
		{
			Src:      "0.31",
			Expected: "0.31",
		},
		{
			Src:      "0.32",
			Expected: "0.32",
		},
		{
			Src:      "0.33",
			Expected: "0.33",
		},
		{
			Src:      "0.34",
			Expected: "0.34",
		},
		{
			Src:      "0.35",
			Expected: "0.35",
		},
		{
			Src:      "0.36",
			Expected: "0.36",
		},
		{
			Src:      "0.37",
			Expected: "0.37",
		},
		{
			Src:      "0.38",
			Expected: "0.38",
		},
		{
			Src:      "0.39",
			Expected: "0.39",
		},
		{
			Src:      "0.40",
			Expected: "0.40",
		},
		{
			Src:      "0.41",
			Expected: "0.41",
		},
		{
			Src:      "0.42",
			Expected: "0.42",
		},
		{
			Src:      "0.43",
			Expected: "0.43",
		},
		{
			Src:      "0.44",
			Expected: "0.44",
		},
		{
			Src:      "0.45",
			Expected: "0.45",
		},
		{
			Src:      "0.46",
			Expected: "0.46",
		},
		{
			Src:      "0.47",
			Expected: "0.47",
		},
		{
			Src:      "0.48",
			Expected: "0.48",
		},
		{
			Src:      "0.49",
			Expected: "0.49",
		},
		{
			Src:      "0.50",
			Expected: "0.50",
		},
		{
			Src:      "0.51",
			Expected: "0.51",
		},
		{
			Src:      "0.52",
			Expected: "0.52",
		},
		{
			Src:      "0.53",
			Expected: "0.53",
		},
		{
			Src:      "0.54",
			Expected: "0.54",
		},
		{
			Src:      "0.55",
			Expected: "0.55",
		},
		{
			Src:      "0.56",
			Expected: "0.56",
		},
		{
			Src:      "0.57",
			Expected: "0.57",
		},
		{
			Src:      "0.58",
			Expected: "0.58",
		},
		{
			Src:      "0.59",
			Expected: "0.59",
		},
		{
			Src:      "0.60",
			Expected: "0.60",
		},
		{
			Src:      "0.61",
			Expected: "0.61",
		},
		{
			Src:      "0.62",
			Expected: "0.62",
		},
		{
			Src:      "0.63",
			Expected: "0.63",
		},
		{
			Src:      "0.64",
			Expected: "0.64",
		},
		{
			Src:      "0.65",
			Expected: "0.65",
		},
		{
			Src:      "0.66",
			Expected: "0.66",
		},
		{
			Src:      "0.67",
			Expected: "0.67",
		},
		{
			Src:      "0.68",
			Expected: "0.68",
		},
		{
			Src:      "0.69",
			Expected: "0.69",
		},
		{
			Src:      "0.70",
			Expected: "0.70",
		},
		{
			Src:      "0.71",
			Expected: "0.71",
		},
		{
			Src:      "0.72",
			Expected: "0.72",
		},
		{
			Src:      "0.73",
			Expected: "0.73",
		},
		{
			Src:      "0.74",
			Expected: "0.74",
		},
		{
			Src:      "0.75",
			Expected: "0.75",
		},
		{
			Src:      "0.76",
			Expected: "0.76",
		},
		{
			Src:      "0.77",
			Expected: "0.77",
		},
		{
			Src:      "0.78",
			Expected: "0.78",
		},
		{
			Src:      "0.79",
			Expected: "0.79",
		},
		{
			Src:      "0.80",
			Expected: "0.80",
		},
		{
			Src:      "0.81",
			Expected: "0.81",
		},
		{
			Src:      "0.82",
			Expected: "0.82",
		},
		{
			Src:      "0.83",
			Expected: "0.83",
		},
		{
			Src:      "0.84",
			Expected: "0.84",
		},
		{
			Src:      "0.85",
			Expected: "0.85",
		},
		{
			Src:      "0.86",
			Expected: "0.86",
		},
		{
			Src:      "0.87",
			Expected: "0.87",
		},
		{
			Src:      "0.88",
			Expected: "0.88",
		},
		{
			Src:      "0.89",
			Expected: "0.89",
		},
		{
			Src:      "0.90",
			Expected: "0.90",
		},
		{
			Src:      "0.91",
			Expected: "0.91",
		},
		{
			Src:      "0.92",
			Expected: "0.92",
		},
		{
			Src:      "0.93",
			Expected: "0.93",
		},
		{
			Src:      "0.94",
			Expected: "0.94",
		},
		{
			Src:      "0.95",
			Expected: "0.95",
		},
		{
			Src:      "0.96",
			Expected: "0.96",
		},
		{
			Src:      "0.97",
			Expected: "0.97",
		},
		{
			Src:      "0.98",
			Expected: "0.98",
		},
		{
			Src:      "0.99",
			Expected: "0.99",
		},
		{
			Src:      "1.00",
			Expected: "1.00",
		},



		{
			Src:      "0.0",
			Expected: "0.0",
		},
		{
			Src:      "0.1",
			Expected: "0.1",
		},
		{
			Src:      "0.2",
			Expected: "0.2",
		},
		{
			Src:      "0.3",
			Expected: "0.3",
		},
		{
			Src:      "0.4",
			Expected: "0.4",
		},
		{
			Src:      "0.5",
			Expected: "0.5",
		},
		{
			Src:      "0.6",
			Expected: "0.6",
		},
		{
			Src:      "0.7",
			Expected: "0.7",
		},
		{
			Src:      "0.8",
			Expected: "0.8",
		},
		{
			Src:      "0.9",
			Expected: "0.9",
		},
		{
			Src:      "1.0",
			Expected: "1.0",
		},



		{
			Src:      "0.12345678901",
			Expected: "0.12345678901",
		},



		{
			Src:      "1",
			Expected: "1",
		},



		{
			Src:      "3.14159265358979323846264338327950",
			Expected: "3.14159265358979323846264338327950",
		},



		{
			Src:      "127",
			Expected: "127",
		},
	}

	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const limit = 20
		for i:=0; i<limit; i++ {
			var builder strings.Builder

			if 0 == randomness.Int63() % 2 {
				builder.WriteRune('-')
			}

			if 0 == randomness.Int63() % 2 {
				switch randomness.Int63() % 10 {
				case 0:
					builder.WriteRune('0')
				case 1:
					fmt.Fprintf(&builder, "%d", randomness.Int63n(10))
				case 2:
					fmt.Fprintf(&builder, "%d", randomness.Int63n(100))
				case 3:
					fmt.Fprintf(&builder, "%d", randomness.Int63n(1000))
				case 4:
					fmt.Fprintf(&builder, "%d", randomness.Int63n(10000))
				case 5:
					fmt.Fprintf(&builder, "%d", randomness.Int63n(100000))
				case 6:
					fmt.Fprintf(&builder, "%d", randomness.Int63n(1000000))
				case 7:
					fmt.Fprintf(&builder, "%d", randomness.Int63n(10000000))
				default:
					fmt.Fprintf(&builder, "%d", randomness.Int63())
				}
			}

			if 0 == randomness.Int63() % 2 {
				builder.WriteRune('.')

				if 0 == randomness.Int63() % 2 {

					switch randomness.Int63() % 10 {
					case 0:
						builder.WriteRune('0')
					case 1:
						fmt.Fprintf(&builder, "%d", randomness.Int63n(10))
					case 2:
						fmt.Fprintf(&builder, "%d", randomness.Int63n(100))
					case 3:
						fmt.Fprintf(&builder, "%d", randomness.Int63n(1000))
					case 4:
						fmt.Fprintf(&builder, "%d", randomness.Int63n(10000))
					case 5:
						fmt.Fprintf(&builder, "%d", randomness.Int63n(100000))
					case 6:
						fmt.Fprintf(&builder, "%d", randomness.Int63n(1000000))
					case 7:
						fmt.Fprintf(&builder, "%d", randomness.Int63n(10000000))
					default:
						fmt.Fprintf(&builder, "%d", randomness.Int63())
					}
				}
			}

			switch builder.String() {
			case "", "-", "-.", ".":
				fmt.Fprintf(&builder, "%d", randomness.Int63())
			}

			str := builder.String()

			test := struct{
				Src      string
				Expected string
			}{
				Src:      str,
				Expected: str,
			}

			tests = append(tests, test)
		}
	}

	{
		moreTests := []struct{
			Src      string
			Expected string
		}{}

		for i, test := range tests {

			// We quote anything of the form .12345 , 12345. , -.12345 , 123,456,789
			if strings.HasPrefix(test.Src, ".") || strings.HasSuffix(test.Src, ".") || strings.HasPrefix(test.Src, "-.") || strings.Contains(test.Src, ",") {
				tests[i].Src = `"`+ test.Src +`"`
				continue
			}

			anotherTest := struct{
				Src      string
				Expected string
			}{
				Src: `"`+ test.Src +`"`,
				Expected: test.Expected,
			}

			moreTests = append(moreTests, anotherTest)
		}

		tests = append(tests, moreTests...)
	}

	for testNumber, test := range tests {


		{
			var actualNullable denary.Nullable

			if expected, actual := denary.Nothing().Nullable(), actualNullable; expected != actual {
				t.Errorf("For test #%d, the denary.Nullable actually had the value %#v, not the expected value of %#v", testNumber, actual, expected)
				continue
			}

			var jsonString string = test.Src

			if err := json.Unmarshal([]byte(jsonString), &actualNullable); nil != err {
				t.Errorf("For test #%d, did not expect an error, for calling json.Unmarshal(), but actually got one: (%T) %q", testNumber, err, err)
				t.Log("SRC:...")
				t.Log(test.Src)
				t.Log("JSON:...")
				t.Log(jsonString)
				continue
			}
			{
				actualValue, err := actualNullable.Return()
				if nil != err {
					t.Errorf("For test #%d, did not expect an error, for calling denary.Nullable.Return(), but actually got one: (%T) %q", testNumber, err, err)
					continue
				}
				if expected, actual := test.Expected, actualValue.String(); expected != actual {
					t.Errorf("For test #%d, the denary.Nullable did not have the value that was actually expected.", testNumber)
					t.Errorf("EXPECTED: «%s»", expected)
					t.Errorf("ACTUAL:   «%s»", actual)
					continue
				}
			}
		}


/*
		var actualNullable denary.Nullable


		result := denary.Parse(src)

		actualValue, ok := result.Unwrap()
		if !ok {
			err := result.Validate()
			t.Errorf("For test #%d, could not unwrap the result (%t), because: (%T) %q", testNumber, ok, err, err)
			t.Logf("SRC: «%s»", src)
			continue
		}

		if expected, actual := test.Expected, actualValue.String(); expected != actual {
			t.Errorf("For test #%d, the denary value that was actually gotten is not what was expected.", testNumber, )
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
*/
	}
}
