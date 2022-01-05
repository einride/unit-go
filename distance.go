package unit

import (
	"fmt"
)

type Distance float64

const meterSymbol = "m"

const (
	Meter Distance = 1
	Mile           = 1 / 0.621371192 * Kilo * Meter
)

func (d Distance) Meters() float64 {
	return float64(d)
}

func (d Distance) Get(as Distance) float64 {
	return float64(d) / float64(as)
}

func (d Distance) String() string {
	return format(float64(d), meterSymbol)
}

func (d *Distance) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		meterSymbol: float64(Meter),
	})
	if err != nil {
		return fmt.Errorf("unmarshal distance: %w", err)
	}
	*d = Distance(parsed)
	return nil
}
