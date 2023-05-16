package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestPressure_Get(t *testing.T) {
	assert.Equal(t, float64(3), (300 * KiloPascal).Get(Bar))
}

func TestPressure_String(t *testing.T) {
	for _, tt := range []struct {
		p   Pressure
		str string
	}{
		{p: 0, str: "0Pa"},
		{p: 2.3 * Kilo * Pascal, str: "2.3kPa"},
		{p: 3 * Milli * Pascal, str: "3mPa"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.p.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s Pressure
				assert.NilError(t, s.UnmarshalString(tt.str))
				assert.Equal(t, tt.p, s)
			})
		})
	}
}
