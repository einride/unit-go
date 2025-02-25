package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestJerk_Get(t *testing.T) {
	assert.Equal(t, float64(3), (3000 * Kilo * Gram).Get(Tonne))
}

func TestJerk_String(t *testing.T) {
	for _, tt := range []struct {
		j   Jerk
		str string
	}{
		{j: 0, str: "0m/s³"},
		{j: 2.3 * Kilo * MeterPerSecondCubed, str: "2.3km/s³"},
		{j: 3 * Milli * MeterPerSecondCubed, str: "3mm/s³"},
	} {
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.j.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s Jerk
				assert.NilError(t, s.UnmarshalString(tt.str))
				assert.Equal(t, tt.j, s)
			})
		})
	}
}
