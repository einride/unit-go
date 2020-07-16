package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestMass_Get(t *testing.T) {
	assert.Equal(t, float64(3), (3000 * Kilo * Gram).Get(Tonne))
}

func TestMass_String(t *testing.T) {
	for _, tt := range []struct {
		m   Mass
		str string
	}{
		{m: 0, str: "0g"},
		{m: 2.3 * Kilo * Gram, str: "2.3kg"},
		{m: 3 * Milli * Gram, str: "3mg"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.m.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s Mass
				assert.NilError(t, s.UnmarshalString(tt.str))
				assert.Equal(t, tt.m, s)
			})
		})
	}
}
