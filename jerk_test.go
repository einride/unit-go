package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJerk_Get(t *testing.T) {
	require.Equal(t, float64(3), (3000 * Kilo * Gram).Get(Tonne))
}

func TestJerk_String(t *testing.T) {
	for _, tt := range []struct {
		j   Jerk
		str string
	}{
		{j: 0, str: "0m/s³"},
		{j: 2.3 * Kilo * MetrePerSecondCubed, str: "2.3km/s³"},
		{j: 3 * Milli * MetrePerSecondCubed, str: "3mm/s³"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				require.Equal(t, tt.str, tt.j.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s Jerk
				require.NoError(t, s.UnmarshalString(tt.str))
				require.Equal(t, tt.j, s)
			})
		})
	}
}
