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

func TestSpeed_String(t *testing.T) {
	for _, tt := range []struct {
		s        Speed
		expected string
	}{
		{s: 0, expected: "0m/s"},
		{s: MetrePerSecond, expected: "1m/s"},
		{s: 2.3 * Centi * MetrePerSecond, expected: "2.3cm/s"},
	} {
		tt := tt
		t.Run(tt.expected, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.s.String())
		})
	}
}
