package unit

import (
	"encoding"
	"fmt"
)

// Speed is the rate of change of an object of its position with time.
type Speed float64

var _ encoding.TextUnmarshaler = (*Speed)(nil)

// MeterPerSecond is the SI unit for measuring Speed.
const MeterPerSecond Speed = 1.0

const meterPerSecondSymbol = "m/s"

// KilometerPerHour is a unit for measuring speed using Kilometer for distance and Hour for time.
const KilometerPerHour = MeterPerSecond / 3.6

const kilometerPerHourSymbol = "km/h"

// MetersPerSecond returns s with the unit of m/s.
func (s Speed) MetersPerSecond() float64 {
	return float64(s)
}

// Get returns s with the unit of as.
func (s Speed) Get(as Speed) float64 {
	return float64(s) / float64(as)
}

// String implements fmt.Stringer.
func (s Speed) String() string {
	return format(float64(s), meterPerSecondSymbol)
}

// UnmarshalString sets *s from str.
func (s *Speed) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		meterPerSecondSymbol:   float64(MeterPerSecond),
		kilometerPerHourSymbol: float64(KilometerPerHour),
	})
	if err != nil {
		return fmt.Errorf("unmarshal speed: %w", err)
	}
	*s = Speed(parsed)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Speed) UnmarshalText(text []byte) error {
	return s.UnmarshalString(string(text))
}
