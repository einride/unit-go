package unit

import (
	"fmt"
)

type Speed float64

const (
	meterPerSecondSymbol   = "m/s"
	kiloMeterPerHourSymbol = "km/h"
)

const (
	MeterPerSecond   Speed = 1.0
	KiloMeterPerHour       = MeterPerSecond / 3.6
)

func (s Speed) MetersPerSecond() float64 {
	return float64(s)
}

func (s Speed) Get(as Speed) float64 {
	return float64(s) / float64(as)
}

func (s Speed) String() string {
	return format(float64(s), meterPerSecondSymbol)
}

func (s *Speed) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		meterPerSecondSymbol:   float64(MeterPerSecond),
		kiloMeterPerHourSymbol: float64(KiloMeterPerHour),
	})
	if err != nil {
		return fmt.Errorf("unmarshal speed: %w", err)
	}
	*s = Speed(parsed)
	return nil
}
