package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestTorque_String(t *testing.T) {
	for _, tt := range []struct {
		t   Torque
		str string
	}{
		{t: 0, str: "0Nm"},
		{t: 2.3 * Kilo * NewtonMeter, str: "2.3kNm"},
		{t: 3 * Milli * NewtonMeter, str: "3mNm"},
	} {
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.t.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s Torque
				assert.NilError(t, s.UnmarshalString(tt.str))
				assert.Equal(t, tt.t, s)
			})
		})
	}
}
