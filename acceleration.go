package unit

import "fmt"

// Acceleration is the rate of change of the velocity of an object with respect to time.
type Acceleration float64

// MeterPerSecondSquared is SI unit for measuring Acceleration.
const MeterPerSecondSquared Acceleration = 1

const meterPerSecondSquaredSymbol = "m/s²"

// MetersPerSecondSquared returns a with the unit of m/s².
func (a Acceleration) MetersPerSecondSquared() float64 {
	return float64(a)
}

// Get returns a with the unit of as.
func (a Acceleration) Get(as Acceleration) float64 {
	return float64(a) / float64(as)
}

// String implements fmt.Stringer.
func (a Acceleration) String() string {
	return format(float64(a), meterPerSecondSquaredSymbol)
}

// UnmarshalString sets *a from s.
func (a *Acceleration) UnmarshalString(s string) error {
	parsed, err := parse(s, map[string]float64{
		meterPerSecondSquaredSymbol: float64(MeterPerSecondSquared),
	})
	if err != nil {
		return fmt.Errorf("unmarshal acceleration: %w", err)
	}
	*a = Acceleration(parsed)
	return nil
}
