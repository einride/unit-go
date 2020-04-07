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

type Angle float64

const (
	Radian Angle = 1.0
	Degree       = Radian / 180 * math.Pi
)

func (a Angle) Radians() float64 {
	return float64(a)
}

func (a Angle) Get(as Angle) float64 {
	return float64(a) / float64(as)
}

func (a Angle) String() string {
	return strconv.FormatFloat(a.Get(Degree), 'f', -1, 64) + degreeSymbol
}

func (a *Angle) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		radianSymbol: float64(Radian),
		degreeSymbol: float64(Degree),
	})
	if err != nil {
		return fmt.Errorf("unmarshal angle: %w", err)
	}
	*a = Angle(parsed)
	return nil
}

// WrapMinusPiPi wraps the current angle in the interval [-pi, pi]
func (a *Angle) WrapMinusPiPi() {
	b := math.Mod(a.Radians()+math.Pi, 2*math.Pi)
	if b < 0 {
		b += 2 * math.Pi
	}
	*a = Angle(b - math.Pi)
}
