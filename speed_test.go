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
		s   Speed
		str string
	}{
		{s: 0, str: "0m/s"},
		{s: MetrePerSecond, str: "1m/s"},
		{s: 2.3 * Centi * MetrePerSecond, str: "2.3cm/s"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				require.Equal(t, tt.str, tt.s.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s Speed
				require.NoError(t, s.UnmarshalString(tt.str))
				require.Equal(t, tt.s, s)
			})
		})
	}
}
