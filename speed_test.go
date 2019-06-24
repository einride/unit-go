package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpeed_Kph(t *testing.T) {
	require.InDelta(t, 3.6, MetrePerSecond.Get(KiloMetrePerHour), 0.0000000000001)
}

func TestSpeed_UnitConversionMPSxKPH(t *testing.T) {
	require.Equal(t, float64(MetrePerSecond), MetrePerSecond.Get(KiloMetrePerHour)*float64(KiloMetrePerHour))
}
