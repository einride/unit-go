package unit

type Distance float64

const (
	Metre Distance = 1
	Mile           = 1 / 0.621371192 * Kilo * Metre
)

func (d Distance) Metres() float64 {
	return float64(d)
}

func (d Distance) Get(as Distance) float64 {
	return float64(d) / float64(as)
}
