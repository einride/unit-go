package unit

import (
	"fmt"
)

// Distance is a numerical measurement of physical length.
type Distance float64

// Meter is the SI unit for measuring Distance.
const Meter Distance = 1

const meterSymbol = "m"

// Mile is a British imperial unit and US customary unit for measuring distance.
const Mile = 1 / 0.621371192 * Kilo * Meter

const mileSymbol = "mi"

// Meters returns d with the unit of m.
func (d Distance) Meters() float64 {
	return float64(d)
}

// Get returns d with the unit of as.
func (d Distance) Get(as Distance) float64 {
	return float64(d) / float64(as)
}

// String implements fmt.Stringer.
func (d Distance) String() string {
	return format(float64(d), meterSymbol)
}

// UnmarshalString sets *d from s.
func (d *Distance) UnmarshalString(s string) error {
	parsed, err := parse(s, map[string]float64{
		meterSymbol: float64(Meter),
		mileSymbol:  float64(Mile),
	})
	if err != nil {
		return fmt.Errorf("unmarshal distance: %w", err)
	}
	*d = Distance(parsed)
	return nil
}
