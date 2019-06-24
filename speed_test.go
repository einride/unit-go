package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpeed_Kph(t *testing.T) {
	require.Equal(t, 3.6, MetrePerSecond.KiloMetresPerHour())
}

func TestSpeed_UnitConversionMPSxKPH(t *testing.T) {
	require.Equal(t, float64(MetrePerSecond), MetrePerSecond.KiloMetresPerHour()*float64(KiloMetrePerHour))
}
