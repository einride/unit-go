package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSharpness_Get(t *testing.T) {
	for _, tt := range []struct {
		msg      string
		d        Sharpness
		as       Sharpness
		expected float64
	}{
		{msg: "rad/m as rad/m", d: 10 * RadianPerMeter, as: RadianPerMeter, expected: 10},
	} {
		t.Run(tt.msg, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.d.Get(tt.as))
		})
	}
}

func TestSharpness_String(t *testing.T) {
	for _, tt := range []struct {
		c   Sharpness
		str string
	}{
		{c: 2.3 * RadianPerMeter, str: "2.3rad/m"},
		{c: 0.1 * RadianPerMeter, str: "0.1rad/m"},
	} {
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.c.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var c Sharpness
				assert.NilError(t, c.UnmarshalString(tt.str))
				assert.Equal(t, tt.c, c)
			})
		})
	}
}
