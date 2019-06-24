package unit

type Acceleration float64

const MetrePerSecondSquared Acceleration = 1

func (a Acceleration) MetresPerSecondSquared() float64 {
	return float64(a)
}

func (a Acceleration) Get(as Acceleration) float64 {
	return float64(a) / float64(as)
}
