package unit

type Distance float64

const (
	milesConstant          = 0.621371192
	Metre         Distance = 1
	Mile                   = 1 / milesConstant * Kilo * Metre
)

func (d Distance) Metres() float64 {
	return float64(d)
}

func (d Distance) Get(as Distance) float64 {
	return float64(d) / float64(as)
}
