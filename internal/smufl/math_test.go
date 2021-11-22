package smufl

import (
	"testing"
)

func TestParseDecimal(t *testing.T) {
	for _, test := range []struct {
		in  string
		out float64
	}{
		{"0.748", 0.748},
		{"0.0", 0.0},
		{"1.1964000000000001", 1.1964},
		{"2.1564444444444444", 2.156444},
	} {
		var z Decimal
		err := z.UnmarshalJSON([]byte(test.in))
		if err != nil {
			t.Error(err)
			continue
		}

		if z.Float64() != test.out {
			t.Errorf("%v != %v", z, test.out)
		}
	}
}

func TestParseDecimalPoint(t *testing.T) {
	for _, test := range []struct {
		in   string
		x, y float64
	}{
		{"[0.748, 0]", 0.748, 0.0},
		{"[0.0, 0.748]", 0.0, 0.748},
		{"[1.1964000000000001, 0]", 1.1964, 0.0},
		{"[2.1564444444444444, 0]", 2.156444, 0.0},
	} {
		var z DecimalPoint
		err := z.UnmarshalJSON([]byte(test.in))
		if err != nil {
			t.Error(err)
			continue
		}

		if z.X.Float64() != test.x {
			t.Errorf("X: %v != %v", z.X, test.x)
		}
		if z.Y.Float64() != test.y {
			t.Errorf("y: %v != %v", z.Y, test.y)
		}
	}
}
