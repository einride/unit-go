package unit

import (
	"fmt"
)

// Torque is the rotational equivalent of linear Force.
type Torque float64

// NewtonMeter is the SI unit for measuring Torque.
const NewtonMeter Torque = 1.0

const newtonMeterSymbol = "Nm"

// NewtonMeters returns t with the unit of Nm.
func (t Torque) NewtonMeters() float64 {
	return float64(t)
}

// Get returns t with the unit of as.
func (t Torque) Get(as Torque) float64 {
	return float64(t) / float64(as)
}

// String implements fmt.Stringer.
func (t Torque) String() string {
	return format(float64(t), newtonMeterSymbol)
}

// UnmarshalString sets *t from s.
func (t *Torque) UnmarshalString(s string) error {
	parsed, err := parse(s, map[string]float64{
		newtonMeterSymbol: float64(NewtonMeter),
	})
	if err != nil {
		return fmt.Errorf("unmarshal torque: %w", err)
	}
	*t = Torque(parsed)
	return nil
}
