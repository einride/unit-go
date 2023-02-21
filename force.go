package unit

import (
	"encoding"
	"fmt"
)

// Force is an influence that can change the motion of an object.
type Force float64

var _ encoding.TextUnmarshaler = (*Force)(nil)

const newtonSymbol = "N"

// Newton is the SI unit for measuring Force.
const Newton Force = 1.0

// Newtons returns f with the unit of N.
func (f Force) Newtons() float64 {
	return float64(f)
}

// Get returns f with the unit of as.
func (f Force) Get(as Force) float64 {
	return float64(f) / float64(as)
}

// String implements fmt.Stringer.
func (f Force) String() string {
	return format(float64(f), newtonSymbol)
}

// UnmarshalString sets *f from s.
func (f *Force) UnmarshalString(s string) error {
	parsed, err := parse(s, map[string]float64{
		newtonSymbol: float64(Newton),
	})
	if err != nil {
		return fmt.Errorf("unmarshal force: %w", err)
	}
	*f = Force(parsed)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (f *Force) UnmarshalText(text []byte) error {
	return f.UnmarshalString(string(text))
}
