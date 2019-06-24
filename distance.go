package unit

type Distance float64

const (
	milesConstant          = 0.621371192
	Metre         Distance = 1
	KiloMetre              = Metre * 1000
	Mile                   = 1 / milesConstant * KiloMetre
)

func (d Distance) Metres() float64 {
	return float64(d)
}

func (d Distance) Miles() float64 {
	return float64(d) * milesConstant / KiloMetre.Metres()
}

func (d Distance) KiloMetres() float64 {
	return float64(d) / 1000
}
