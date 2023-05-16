package unit

import (
	"encoding"
	"fmt"
)

// Pressure is the force applied perpendicular to the surface of an object per unit area over which that
// force is distributed.
type Pressure float64

var _ encoding.TextUnmarshaler = (*Pressure)(nil)

// Pascal is the SI unit for measuring Pressure.
const Pascal Pressure = 1

const pascalSymbol = "Pa"

// KiloPascal is 1000 times the SI unit Pascal.
const KiloPascal Pressure = 1000 * Pascal

const kiloPascalSymbol = "kPa"

// KiloPascal returns p with the unit of kPa.
func (p Pressure) KiloPascal() float64 {
	return p.Get(KiloPascal)
}

// Bar is 100'000 times the SI unit Pascal.
const Bar Pressure = 1e5 * Pascal

const barSymbol = "bar"

// Bar returns p with the unit of 1 bar.
func (p Pressure) Bar() float64 {
	return p.Get(Bar)
}

// Get returns p with the unit of as.
func (p Pressure) Get(as Pressure) float64 {
	return float64(p) / float64(as)
}

// String implements fmt.Stringer.
func (p Pressure) String() string {
	return format(float64(p), pascalSymbol)
}

// UnmarshalString sets *p from s.
func (p *Pressure) UnmarshalString(s string) error {
	parsed, err := parse(s, map[string]float64{
		pascalSymbol:     float64(Pascal),
		kiloPascalSymbol: float64(KiloPascal),
		barSymbol:        float64(Bar),
	})
	if err != nil {
		return fmt.Errorf("unmarshal pressure: %w", err)
	}
	*p = Pressure(parsed)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (p *Pressure) UnmarshalText(text []byte) error {
	return p.UnmarshalString(string(text))
}
