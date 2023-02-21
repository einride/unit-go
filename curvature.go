package unit

import (
	"encoding"
	"fmt"
	"strconv"
	"strings"
)

// Curvature is a numerical measurement of inverse physical length.
type Curvature float64

var _ encoding.TextUnmarshaler = (*Curvature)(nil)

// PerMeter is the SI unit for measuring Curvature.
const PerMeter Curvature = 1

const perMeterSymbol = "/m"

// PerMeters returns c with the unit of 1/m.
func (c Curvature) PerMeters() float64 {
	return float64(c)
}

// Get returns c with the unit of as.
func (c Curvature) Get(as Curvature) float64 {
	return float64(c) / float64(as)
}

// String implements fmt.Stringer.
func (c Curvature) String() string {
	return format(float64(c), perMeterSymbol)
}

// UnmarshalString sets *c from s.
func (c *Curvature) UnmarshalString(s string) error {
	if !strings.HasSuffix(s, perMeterSymbol) {
		return fmt.Errorf("unmarshal curvature: input '%s' doesn't have the '%s' suffix", s, perMeterSymbol)
	}
	sNoUnit := strings.TrimSuffix(s, perMeterSymbol)
	if sNoUnit != strings.TrimSpace(sNoUnit) {
		return fmt.Errorf("unmarshal curvature: input '%s' has invalid whitespace", s)
	}
	magnitude, err := strconv.ParseFloat(sNoUnit, 64)
	if err != nil {
		return fmt.Errorf("unmarshal curvature: parse '%s': %w", s, err)
	}
	*c = Curvature(magnitude)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (c *Curvature) UnmarshalText(text []byte) error {
	return c.UnmarshalString(string(text))
}
