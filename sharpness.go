package unit

import (
	"fmt"
	"strconv"
	"strings"
)

// Sharpness is a numerical measurement of angular change per physical length.
// And example is the sharpness of a curve.
type Sharpness float64

// RadianPerMeter is the SI unit for measuring Sharpness.
const RadianPerMeter Sharpness = 1

const radianPerMeterSymol = "rad/m"

// RadianPerMeter returns sharpness as a float64 with the unit of rad/m.
func (s Sharpness) RadianPerMeter() float64 {
	return float64(s)
}

// Get returns the sharpness with the unit of as.
func (s Sharpness) Get(as Sharpness) float64 {
	return float64(s) / float64(as)
}

// String implements fmt.Stringer.
func (s Sharpness) String() string {
	return format(float64(s), radianPerMeterSymol)
}

// UnmarshalString sets *c from s.
func (s *Sharpness) UnmarshalString(str string) error {
	if !strings.HasSuffix(str, radianPerMeterSymol) {
		return fmt.Errorf("unmarshal sharpness: input '%s' doesn't have the '%s' suffix", s, radianPerMeterSymol)
	}
	strNoUnit := strings.TrimSuffix(str, radianPerMeterSymol)
	if strNoUnit != strings.TrimSpace(strNoUnit) {
		return fmt.Errorf("unmarshal sharpness: input '%s' has invalid whitespace", s)
	}
	magnitude, err := strconv.ParseFloat(strNoUnit, 64)
	if err != nil {
		return fmt.Errorf("unmarshal sharpness: parse '%s': %w", s, err)
	}
	*s = Sharpness(magnitude)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Sharpness) UnmarshalText(text []byte) error {
	return s.UnmarshalString(string(text))
}
