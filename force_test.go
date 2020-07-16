package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestForce_String(t *testing.T) {
	for _, tt := range []struct {
		f   Force
		str string
	}{
		{f: 0, str: "0N"},
		{f: 2.3 * Kilo * Newton, str: "2.3kN"},
		{f: 3 * Milli * Newton, str: "3mN"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.f.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s Force
				assert.NilError(t, s.UnmarshalString(tt.str))
				assert.Equal(t, tt.f, s)
			})
		})
	}
}
