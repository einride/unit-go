package unit

import "fmt"

type Force float64

const newtonSymbol = "N"

const Newton Force = 1.0

func (t Force) Newtons() float64 {
	return float64(t)
}

func (t Force) Get(as Force) float64 {
	return float64(t) / float64(as)
}

func (t Force) String() string {
	return format(float64(t), newtonSymbol)
}

func (t *Force) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		newtonSymbol: float64(Newton),
	})
	if err != nil {
		return fmt.Errorf("unmarshal force: %w", err)
	}
	*t = Force(parsed)
	return nil
}
