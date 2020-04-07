package unit

import "fmt"

type Jerk float64

const metresPerSecondCubedSymbol = "m/s³"

const MetrePerSecondCubed Jerk = 1.0

func (m Jerk) MetresPerSecondCubed() float64 {
	return float64(m)
}

func (m Jerk) Get(as Jerk) float64 {
	return float64(m) / float64(as)
}

func (m Jerk) String() string {
	return format(float64(m), metresPerSecondCubedSymbol)
}

func (m *Jerk) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		metresPerSecondCubedSymbol: float64(MetrePerSecondCubed),
	})
	if err != nil {
		return fmt.Errorf("unmarshal jerk: %w", err)
	}
	*m = Jerk(parsed)
	return nil
}
