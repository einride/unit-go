package unit

import "math"

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
