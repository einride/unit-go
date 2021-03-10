package unit

import (
	"fmt"
)

type (
	Speed            float64
	KilometerPerHour float64
	MeterPerSecond   float64
)

const (
	metrePerSecondSymbol   = "m/s"
	kiloMetrePerHourSymbol = "km/h"
)

const (
	mpsToKph = 3.6
	kphToMps = 1 / mpsToKph
)

func (m MeterPerSecond) MetresPerSecond() float64 {
	return float64(m)
}

func (m MeterPerSecond) AsKilometerPerHour() KilometerPerHour {
	return KilometerPerHour(m * mpsToKph)
}

func (m MeterPerSecond) String() string {
	return format(float64(m), metrePerSecondSymbol)
}

func (m *MeterPerSecond) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		metrePerSecondSymbol:   float64(1),
		kiloMetrePerHourSymbol: kphToMps,
	})
	if err != nil {
		return fmt.Errorf("unmarshal speed: %w", err)
	}
	*m = MeterPerSecond(parsed)
	return nil
}

func (k KilometerPerHour) AsMetresPerSecond() MeterPerSecond {
	return MeterPerSecond(k * kphToMps)
}

func (k KilometerPerHour) KilometerPerHour() float64 {
	return float64(k)
}

func (k KilometerPerHour) String() string {
	return format(float64(k), kiloMetrePerHourSymbol)
}

func (k *KilometerPerHour) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		metrePerSecondSymbol:   mpsToKph,
		kiloMetrePerHourSymbol: float64(1),
	})
	if err != nil {
		return fmt.Errorf("unmarshal speed: %w", err)
	}
	*k = KilometerPerHour(parsed)
	return nil
}
