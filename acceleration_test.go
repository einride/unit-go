package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAcceleration_String(t *testing.T) {
	require.Equal(t, "1.2345m/s²", (123.45 * Centi * MetrePerSecondSquared).String())
}
