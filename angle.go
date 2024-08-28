package unit

import (
	"encoding"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

// Angle is the measure of a plane angle or rotation.
type Angle float64

var _ encoding.TextUnmarshaler = (*Angle)(nil)

// Radian is the SI unit for measuring an Angle.
const Radian Angle = 1.0

const radianSymbol = "rad"

// Degree is a measurement of a plane angle in which one full rotation is 360 degrees.
const Degree = Radian / 180 * math.Pi

const degreeSymbol = "°"

const degreeText = "deg"

// Radians returns the angle with the unit of radians.
func (a Angle) Radians() float64 {
	return float64(a)
}

// Degrees returns the angle with the unit of degrees.
func (a Angle) Degrees() float64 {
	return float64(a * 180 / math.Pi)
}

// FromRadians returns an Angle from radians as float64.
func FromRadians(a float64) Angle {
	return Angle(a)
}

// FromDegrees returns an Angle from degrees as float64.
func FromDegrees(a float64) Angle {
	return Angle(a) * Degree
}

// WrapMinusPiPi wraps the current angle in the interval [-pi, pi[.
func (a Angle) WrapMinusPiPi() Angle {
	b := math.Mod(a.Radians()+math.Pi, 2*math.Pi)
	if b < 0 {
		b += 2 * math.Pi
	}
	return Angle(b - math.Pi)
}

// WrapZeroTwoPi wraps the current angle in the interval [0, 2*pi[.
func (a Angle) WrapZeroTwoPi() Angle {
	b := math.Mod(a.Radians(), 2*math.Pi)
	if b < 0 {
		b += 2 * math.Pi
	}
	return Angle(b)
}

// Get returns a with the unit of as.
func (a Angle) Get(as Angle) float64 {
	return float64(a) / float64(as)
}

// String implements fmt.Stringer.
func (a Angle) String() string {
	return strconv.FormatFloat(a.Get(Degree), 'f', -1, 64) + degreeSymbol
}

// UnmarshalString sets *a from s.
func (a *Angle) UnmarshalString(s string) error {
	parsed, err := parse(s, map[string]float64{
		radianSymbol: float64(Radian),
		degreeSymbol: float64(Degree),
		degreeText:   float64(Degree),
	})
	if err != nil {
		return fmt.Errorf("unmarshal angle: %w", err)
	}
	*a = Angle(parsed)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (a *Angle) UnmarshalText(text []byte) error {
	return a.UnmarshalString(string(text))
}

// UnmarshalJSON implements JSON unmarshalling for the Angle type.
// The type is represented as radians in JSON.
func (a *Angle) UnmarshalJSON(data []byte) error {
	var angle float64
	err := json.Unmarshal(data, &angle)
	if err != nil {
		return err
	}
	*a = FromRadians(angle)
	return nil
}

func (a *Angle) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Radians())
}
