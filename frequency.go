package unit

import "fmt"

type Frequency float64

const hertzSymbol = "Hz"

const Hertz Frequency = 1.0

func (f Frequency) Hertz() float64 {
	return float64(f)
}

func (f Frequency) Get(as Frequency) float64 {
	return float64(f) / float64(as)
}

func (f Frequency) String() string {
	return format(float64(f), hertzSymbol)
}

func (f *Frequency) UnmarshalString(str string) error {
	parsed, err := parse(str, map[string]float64{
		hertzSymbol: float64(Hertz),
	})
	if err != nil {
		return fmt.Errorf("unmarshal frequency: %w", err)
	}
	*f = Frequency(parsed)
	return nil
}
