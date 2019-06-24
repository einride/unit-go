package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistance(t *testing.T) {
	require.Equal(t, 1000., (Kilo * Metre).Metres())
	require.Equal(t, 1., Metre.Metres())
}

func TestDistance_KiloMetresPerHour(t *testing.T) {
	require.Equal(t, 1., (Kilo * Metre).Get(Kilo*Metre))
}

func TestDistance_Mile(t *testing.T) {
	require.InDelta(t, 1., Mile.Get(Mile), 0.0000000000001)
}

func TestDistance_String(t *testing.T) {
	for _, tt := range []struct {
		d        Distance
		expected string
	}{
		{d: Centi * Metre, expected: "1cm"},
		{d: Kilo * Metre, expected: "1km"},
		{d: 2.3 * Kilo * Metre, expected: "2.3km"},
	} {
		tt := tt
		t.Run(tt.expected, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.d.String())
		})
	}
}
