package unit

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAngle_FromDegrees(t *testing.T) {
	require.Equal(t, math.Pi*Radian, 180*Degree)
}

func TestAngle_ToDegrees(t *testing.T) {
	require.Equal(t, 180.0, (Radian * math.Pi).Get(Degree))
}
