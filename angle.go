package unit

import (
	"fmt"
	"math"
	"strconv"
)

const (
	radianSymbol = "rad"
	degreeSymbol = "°"
)

type (
	Radians float64
	Degree  float64
)

const (
	degreeToRadian = math.Pi / 180
	radianToDegree = 1 / degreeToRadian
)

func (r Radians) Radians() float64 {
	return float64(r)
}

func (r Radians) AsDegree() Degree {
	return Degree(r * radianToDegree)
}

func (r Radians) String() string {
	return strconv.FormatFloat(float64(r.AsDegree()), 'f', -1, 64) + degreeSymbol
}

func (r *Radians) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		radianSymbol: float64(1),
		degreeSymbol: degreeToRadian,
	})
	if err != nil {
		return fmt.Errorf("unmarshal angle: %w", err)
	}
	*r = Radians(parsed)
	return nil
}

// WrapMinusPiPi wraps the current angle in the interval [-pi, pi]
func (r *Radians) WrapMinusPiPi() Radians {
	b := math.Mod(r.Radians()+math.Pi, 2*math.Pi)
	if b < 0 {
		b += 2 * math.Pi
	}
	return Radians(b - math.Pi)
}

func (d Degree) Degrees() float64 {
	return float64(d)
}

func (d Degree) AsRadians() Radians {
	return Radians(d * degreeToRadian)
}

func (d Degree) String() string {
	return strconv.FormatFloat(d.Degrees(), 'f', -1, 64) + degreeSymbol
}

func (d *Degree) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		degreeSymbol: float64(1),
		radianSymbol: radianToDegree,
	})
	if err != nil {
		return fmt.Errorf("unmarshal angle: %w", err)
	}
	*d = Degree(parsed)
	return nil
}
