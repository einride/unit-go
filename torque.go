package unit

import "golang.org/x/xerrors"

type Torque float64

const newtonMetreSymbol = "Nm"

const NewtonMetre Torque = 1.0

func (t Torque) NewtonMetres() float64 {
	return float64(t)
}

func (t Torque) Get(as Torque) float64 {
	return float64(t) / float64(as)
}

func (t Torque) String() string {
	return format(float64(t), newtonMetreSymbol)
}

func (t *Torque) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		newtonMetreSymbol: float64(NewtonMetre),
	})
	if err != nil {
		return xerrors.Errorf("unmarshal torque: %w", err)
	}
	*t = Torque(parsed)
	return nil
}
