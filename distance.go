package unit

import "golang.org/x/xerrors"

type Distance float64

const metreSymbol = "m"

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

func (d Distance) String() string {
	return format(float64(d), metreSymbol)
}

func (d *Distance) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		metreSymbol: float64(Metre),
	})
	if err != nil {
		return xerrors.Errorf("unmarshal distance: %w", err)
	}
	*d = Distance(parsed)
	return nil
}
