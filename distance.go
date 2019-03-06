package unit

type Distance float64

const (
	milesConstant          = 0.621371192
	Meter         Distance = 1
	Kilometer              = Meter * 1000
	Mile                   = 1 / milesConstant * Kilometer
)

func (d Distance) Meters() float64 {
	return float64(d)
}

func (d Distance) Miles() float64 {
	return float64(d) * milesConstant / Kilometer.Meters()
}

func (d Distance) Kilometers() float64 {
	return float64(d) / 1000
}
