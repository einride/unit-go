package unit

import (
	"fmt"
)

type Torque float64

const newtonMeterSymbol = "Nm"

const NewtonMeter Torque = 1.0

func (t Torque) NewtonMeters() float64 {
	return float64(t)
}

func (t Torque) Get(as Torque) float64 {
	return float64(t) / float64(as)
}

func (t Torque) String() string {
	return format(float64(t), newtonMeterSymbol)
}

func (t *Torque) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		newtonMeterSymbol: float64(NewtonMeter),
	})
	if err != nil {
		return fmt.Errorf("unmarshal torque: %w", err)
	}
	*t = Torque(parsed)
	return nil
}
