package smufl

import (
	"encoding/json"
	"fmt"
	"strconv"

	"golang.org/x/image/math/fixed"
)

// Decimal is a decimal represented as v * 1e6, with the lowest
// bit saying whether the value has been set or not.
type Decimal int64

func ParseFloat(v string) (Decimal, error) {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %q: %w", v, err)
	}
	// the Decimal values in json are reperesented as staff spaces
	return Decimal(f*1e6/4)<<1 | 1, nil
}

func (v Decimal) IsUnset() bool { return v == 0 }

func (v *Decimal) UnmarshalJSON(xs []byte) error {
	d, err := ParseFloat(string(xs))
	if err != nil {
		return err
	}
	*v = d
	return nil
}

func (v Decimal) Px(ppem fixed.Int26_6) fixed.Int26_6 {
	// TODO: verify correctness and implement without overflow.
	return fixed.Int26_6(Decimal(ppem) * (v >> 1) / 1e6)
}

func (v Decimal) Float64() float64 { return float64(v>>1) / 1e6 }
func (v Decimal) String() string   { return fmt.Sprintf("%f", v.Float64()) }

type DecimalPoint struct{ X, Y Decimal }

func (v *DecimalPoint) UnmarshalJSON(xs []byte) error {
	var q [2]json.Number
	if err := json.Unmarshal(xs, &q); err != nil {
		return err
	}
	if err := v.X.UnmarshalJSON([]byte(q[0].String())); err != nil {
		return err
	}
	if err := v.Y.UnmarshalJSON([]byte(q[1].String())); err != nil {
		return err
	}
	return nil
}

func (v DecimalPoint) String() string {
	return fmt.Sprintf("[%v, %v]", v.X, v.Y)
}
