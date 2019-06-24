package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistance(t *testing.T) {
	require.InDelta(t, 1000/milesConstant, (1 * Mile).Metres(), 0.0000000001)
	require.Equal(t, 1000., 1*KiloMetre.Metres())
	require.Equal(t, 1., 1*Metre.Metres())
}

func TestSpeed_KiloMetre(t *testing.T) {
	require.Equal(t, 1., (1 * KiloMetre).KiloMetres())
}

func TestSpeed_Mile(t *testing.T) {
	require.InDelta(t, 1., (1 * Mile).Miles(), 0.0000000000001)
}
