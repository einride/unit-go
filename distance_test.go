package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistance(t *testing.T) {
	require.InDelta(t, 1000/milesConstant, (1 * Mile).Meters(), 0.0000000001)
	require.Equal(t, 1000., 1*Kilometer.Meters())
	require.Equal(t, 1., 1*Meter.Meters())
}

func TestSpeed_Kilometer(t *testing.T) {
	require.Equal(t, 1., (1 * Kilometer).Kilometers())
}

func TestSpeed_Mile(t *testing.T) {
	require.InDelta(t, 1., (1 * Mile).Miles(), 0.0000000000001)
}
