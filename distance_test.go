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

func TestSpeed_Mile(t *testing.T) {
	require.InDelta(t, 1., Mile.Get(Mile), 0.0000000000001)
}
