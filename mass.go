package unit

import (
	"encoding"
	"fmt"
)

// Mass is the quantity of matter in a physical body.
type Mass float64

var _ encoding.TextUnmarshaler = (*Mass)(nil)

// Kilogram is the SI unit for measuring Mass.
const Kilogram = Kilo * Gram

const kilogramSymbol = "kg"

// Gram is one thousandth of the SI unit for measuring mass, the Kilogram.
const Gram Mass = 1e-3

const gramSymbol = "g"

// Tonne is a thousand of the SI unit for measuring mass, the Kilogram.
const Tonne = 1e6 * Gram

const tonneSymbol = "t"

// Kilograms returns m with the unit of kg.
func (m Mass) Kilograms() float64 {
	return m.Get(Kilogram)
}

// Get returns m with the unit of as.
func (m Mass) Get(as Mass) float64 {
	return float64(m) / float64(as)
}

// String implements fmt.Stringer.
func (m Mass) String() string {
	return format(float64(m), gramSymbol)
}

// UnmarshalString sets *m from s.
func (m *Mass) UnmarshalString(s string) error {
	parsed, err := parse(s, map[string]float64{
		gramSymbol:     float64(Gram),
		kilogramSymbol: float64(Kilogram),
		tonneSymbol:    float64(Tonne),
	})
	if err != nil {
		return fmt.Errorf("unmarshal mass: %w", err)
	}
	*m = Mass(parsed)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (m *Mass) UnmarshalText(text []byte) error {
	return m.UnmarshalString(string(text))
}
