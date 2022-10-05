package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestCurvature_Get(t *testing.T) {
	for _, tt := range []struct {
		msg      string
		d        Curvature
		as       Curvature
		expected float64
	}{
		{msg: "1/m as 1/m", d: 10 * PerMeter, as: PerMeter, expected: 10},
	} {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.d.Get(tt.as))
		})
	}
}

func TestCurvature_String(t *testing.T) {
	for _, tt := range []struct {
		c   Curvature
		str string
	}{
		{c: 2.3 * PerMeter, str: "2.3/m"},
		{c: 0.1 * PerMeter, str: "0.1/m"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.c.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var c Curvature
				assert.NilError(t, c.UnmarshalString(tt.str))
				assert.Equal(t, tt.c, c)
			})
		})
	}
}
