package unit

import "fmt"

const meterPerSecondSquaredSymbol = "m/s²"

type Acceleration float64

const MeterPerSecondSquared Acceleration = 1

func (a Acceleration) MetersPerSecondSquared() float64 {
	return float64(a)
}

func (a Acceleration) Get(as Acceleration) float64 {
	return float64(a) / float64(as)
}

func (a Acceleration) String() string {
	return format(float64(a), meterPerSecondSquaredSymbol)
}

func (a *Acceleration) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		meterPerSecondSquaredSymbol: float64(MeterPerSecondSquared),
	})
	if err != nil {
		return fmt.Errorf("unmarshal acceleration: %w", err)
	}
	*a = Acceleration(parsed)
	return nil
}
