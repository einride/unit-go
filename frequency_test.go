package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestFrequency_String(t *testing.T) {
	for _, tt := range []struct {
		f   Frequency
		str string
	}{
		{f: 0, str: "0Hz"},
		{f: 2.3 * Kilo * Hertz, str: "2.3kHz"},
		{f: 3 * Milli * Hertz, str: "3mHz"},
	} {
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.f.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s Frequency
				assert.NilError(t, s.UnmarshalString(tt.str))
				assert.Equal(t, tt.f, s)
			})
		})
	}
}
