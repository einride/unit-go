package unit

type Speed float64

const (
	MetrePerSecond   Speed = 1.0
	KiloMetrePerHour       = MetrePerSecond / 3.6
)

func (s Speed) MetersPerSecond() float64 {
	return float64(s)
}

func (s Speed) KiloMetresPerHour() float64 {
	return float64(s) * 3.6
}
