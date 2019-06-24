package unit

type Speed float64

const metrePerSecondSymbol = "m/s"

const (
	MetrePerSecond   Speed = 1.0
	KiloMetrePerHour       = MetrePerSecond / 3.6
)

func (s Speed) MetresPerSecond() float64 {
	return float64(s)
}

func (s Speed) Get(as Speed) float64 {
	return float64(s) / float64(as)
}

func (s Speed) String() string {
	return formatWithPrefixAndSymbol(float64(s), metrePerSecondSymbol)
}
