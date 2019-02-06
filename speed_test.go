package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpeed_Kph(t *testing.T) {
	require.Equal(t, 3.6, MeterPerSecond.KilometersPerHour())
}

func TestSpeed_UnitConversionMPSxKPH(t *testing.T) {
	require.Equal(t, float64(MeterPerSecond), MeterPerSecond.KilometersPerHour()*float64(KilometerPerHour))
}
