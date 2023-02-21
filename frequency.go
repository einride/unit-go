package unit

import (
	"encoding"
	"fmt"
)

// Frequency is the number of occurrences of a repeating event per unit of time.
type Frequency float64

var _ encoding.TextUnmarshaler = (*Frequency)(nil)

// Hertz is the SI unit for measuring Frequency.
const Hertz Frequency = 1.0

const hertzSymbol = "Hz"

// Hertz returns f with the unit of Hz.
func (f Frequency) Hertz() float64 {
	return float64(f)
}

// Get returns f with the unit of as.
func (f Frequency) Get(as Frequency) float64 {
	return float64(f) / float64(as)
}

// String implements fmt.Stringer.
func (f Frequency) String() string {
	return format(float64(f), hertzSymbol)
}

// UnmarshalString sets *f from s.
func (f *Frequency) UnmarshalString(s string) error {
	parsed, err := parse(s, map[string]float64{
		hertzSymbol: float64(Hertz),
	})
	if err != nil {
		return fmt.Errorf("unmarshal frequency: %w", err)
	}
	*f = Frequency(parsed)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (f *Frequency) UnmarshalText(text []byte) error {
	return f.UnmarshalString(string(text))
}
