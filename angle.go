package unit

import (
	"encoding"
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

// WrapMinusPiPi wraps the current angle in the interval [-pi, pi].
func (a Angle) WrapMinusPiPi() Angle {
	const limitForModulo = 5 * math.Pi

	if math.Abs(a.Radians()) > limitForModulo {
		return a.wrapMinusPiPiModulo()
	}
	return a.wrapMinusPiPiFast()
}

// wrapMinusPiPiModulo wraps the current angle in the interval [-pi, pi] using
// modulo operations, faster for huge angle values.
func (a Angle) wrapMinusPiPiModulo() Angle {
	b := math.Mod(a.Radians()+math.Pi, 2*math.Pi)
	if b < 0 {
		b += 2 * math.Pi
	}
	return Angle(b - math.Pi)
}

// wrapMinusPiPiFast wraps the current angle in the interval [-pi, pi] using for
// loops, faster for small angle values.
func (a Angle) wrapMinusPiPiFast() Angle {
	const twoPi = 2 * math.Pi
	for a < -math.Pi {
		a += twoPi
	}
	for a >= math.Pi {
		a -= twoPi
	}
	return a
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
