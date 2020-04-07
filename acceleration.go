package unit

import "fmt"

const metrePerSecondSquaredSymbol = "m/s²"

type Acceleration float64

const MetrePerSecondSquared Acceleration = 1

func (a Acceleration) MetresPerSecondSquared() float64 {
	return float64(a)
}

func (a Acceleration) Get(as Acceleration) float64 {
	return float64(a) / float64(as)
}

func (a Acceleration) String() string {
	return format(float64(a), metrePerSecondSquaredSymbol)
}

func (a *Acceleration) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		metrePerSecondSquaredSymbol: float64(MetrePerSecondSquared),
	})
	if err != nil {
		return fmt.Errorf("unmarshal acceleration: %w", err)
	}
	*a = Acceleration(parsed)
	return nil
}
