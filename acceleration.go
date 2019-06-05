package unit

type Acceleration float64

const (
	MetersPerSecondSquared Acceleration = 1
)

func (acceleration Acceleration) MetersPerSecondSquared() float64 {
	return float64(acceleration)
}
