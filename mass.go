package unit

import "golang.org/x/xerrors"

type Mass float64

const (
	gramSymbol  = "g"
	tonneSymbol = "t"
)

const (
	Gram     Mass = 1.0
	KiloGram      = Kilo * Gram
	Tonne         = 1e6 * Gram
)

func (m Mass) Grams() float64 {
	return float64(m)
}

func (m Mass) KiloGrams() float64 {
	return m.Get(Kilo * Gram)
}

func (m Mass) Get(as Mass) float64 {
	return float64(m) / float64(as)
}

func (m Mass) String() string {
	return format(float64(m), gramSymbol)
}

func (m *Mass) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		gramSymbol:  float64(Gram),
		tonneSymbol: float64(Tonne),
	})
	if err != nil {
		return xerrors.Errorf("unmarshal mass: %w", err)
	}
	*m = Mass(parsed)
	return nil
}
