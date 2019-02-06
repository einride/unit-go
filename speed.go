package unit

type Speed float64

const (
	MeterPerSecond   Speed = 1.0
	KilometerPerHour       = MeterPerSecond / 3.6
)

func (s Speed) MetersPerSecond() float64 {
	return float64(s)
}

func (s Speed) KilometersPerHour() float64 {
	return float64(s) * 3.6
}
