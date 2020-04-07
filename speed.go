package unit

import (
	"fmt"
)

type Speed float64

const (
	metrePerSecondSymbol   = "m/s"
	kiloMetrePerHourSymbol = "km/h"
)

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
	return format(float64(s), metrePerSecondSymbol)
}

func (s *Speed) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		metrePerSecondSymbol:   float64(MetrePerSecond),
		kiloMetrePerHourSymbol: float64(KiloMetrePerHour),
	})
	if err != nil {
		return fmt.Errorf("unmarshal speed: %w", err)
	}
	*s = Speed(parsed)
	return nil
}
