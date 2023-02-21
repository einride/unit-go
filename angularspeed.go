package unit

import (
	"encoding"
	"fmt"
	"math"
	"strconv"
)

// AngularSpeed is a measure of rotation rate.
type AngularSpeed float64

var _ encoding.TextUnmarshaler = (*AngularSpeed)(nil)

// RadianPerSecond is the SI unit for measuring AngularSpeed.
const RadianPerSecond AngularSpeed = 1.0

const radianPerSecondSymbol = "rad/s"

// RPM is the number of revolutions in one minute.
const RPM = RadianPerSecond * (2 * math.Pi) / 60

const rpmSymbol = "RPM"

// DegreePerSecond is angular speed measured in degrees per second.
const DegreePerSecond = RadianPerSecond * (180 / math.Pi)

const degreePerSecondSymbol = "°/s"

// RadiansPerSecond returns a with the unit of rad/s.
func (a AngularSpeed) RadiansPerSecond() float64 {
	return float64(a)
}

// Get returns a with the unit of as.
func (a AngularSpeed) Get(as AngularSpeed) float64 {
	return float64(a) / float64(as)
}

// String implements fmt.Stringer.
func (a AngularSpeed) String() string {
	return strconv.FormatFloat(float64(a), 'f', -1, 64) + radianPerSecondSymbol
}

// UnmarshalString sets *a from s.
func (a *AngularSpeed) UnmarshalString(s string) error {
	parsed, err := parse(s, map[string]float64{
		radianPerSecondSymbol: float64(RadianPerSecond),
		rpmSymbol:             float64(RPM),
		degreePerSecondSymbol: float64(DegreePerSecond),
	})
	if err != nil {
		return fmt.Errorf("unmarshal angular speed: %w", err)
	}
	*a = AngularSpeed(parsed)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (a *AngularSpeed) UnmarshalText(text []byte) error {
	return a.UnmarshalString(string(text))
}
