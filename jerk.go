package unit

import "fmt"

type Jerk float64

const metersPerSecondCubedSymbol = "m/s³"

const MeterPerSecondCubed Jerk = 1.0

func (m Jerk) MetersPerSecondCubed() float64 {
	return float64(m)
}

func (m Jerk) Get(as Jerk) float64 {
	return float64(m) / float64(as)
}

func (m Jerk) String() string {
	return format(float64(m), metersPerSecondCubedSymbol)
}

func (m *Jerk) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		metersPerSecondCubedSymbol: float64(MeterPerSecondCubed),
	})
	if err != nil {
		return fmt.Errorf("unmarshal jerk: %w", err)
	}
	*m = Jerk(parsed)
	return nil
}
