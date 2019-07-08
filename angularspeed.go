package unit

import (
	"math"
	"strconv"

	"golang.org/x/xerrors"
)

const radianPerSecondSymbol = "rad/s"

type AngularSpeed float64

const (
	RadianPerSecond AngularSpeed = 1.0
	RPM                          = RadianPerSecond * (2 * math.Pi) / 60
)

func (a AngularSpeed) RadiansPerSecond() float64 {
	return float64(a)
}

func (a AngularSpeed) Get(as AngularSpeed) float64 {
	return float64(a) / float64(as)
}

func (a AngularSpeed) String() string {
	return strconv.FormatFloat(float64(a), 'f', -1, 64) + radianPerSecondSymbol
}

func (a *AngularSpeed) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		radianPerSecondSymbol: float64(RadianPerSecond),
	})
	if err != nil {
		return xerrors.Errorf("unmarshal angular speed: %w", err)
	}
	*a = AngularSpeed(parsed)
	return nil
}
