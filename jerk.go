package unit

import "fmt"

// Jerk is the rate at which an object's acceleration changes with respect to time.
type Jerk float64

// MeterPerSecondCubed is the SI unit for measuring Jerk.
const MeterPerSecondCubed Jerk = 1.0

const metersPerSecondCubedSymbol = "m/s³"

// MetersPerSecondCubed returns j with the unit of m/s³.
func (j Jerk) MetersPerSecondCubed() float64 {
	return float64(j)
}

// Get returns j with the unit of as.
func (j Jerk) Get(as Jerk) float64 {
	return float64(j) / float64(as)
}

// String implements fmt.Stringer.
func (j Jerk) String() string {
	return format(float64(j), metersPerSecondCubedSymbol)
}

// UnmarshalString sets *j from s.
func (j *Jerk) UnmarshalString(s string) error {
	parsed, err := parse(s, map[string]float64{
		metersPerSecondCubedSymbol: float64(MeterPerSecondCubed),
	})
	if err != nil {
		return fmt.Errorf("unmarshal jerk: %w", err)
	}
	*j = Jerk(parsed)
	return nil
}
