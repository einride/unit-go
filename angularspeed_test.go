package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestAngularSpeed_Get(t *testing.T) {
	assert.Equal(t, float64(9.549296585513721), (RadianPerSecond).Get(RPM))
	assert.Equal(t, float64(0.017453292519943295), RadianPerSecond.Get(DegreePerSecond))
}

func TestAngularSpeed_String(t *testing.T) {
	for _, tt := range []struct {
		a   AngularSpeed
		str string
	}{
		{a: 0, str: "0rad/s"},
		{a: 2.3 * RadianPerSecond, str: "2.3rad/s"},
		{a: 3 * Milli * RadianPerSecond, str: "0.003rad/s"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.a.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s AngularSpeed
				assert.NilError(t, s.UnmarshalString(tt.str))
				assert.Equal(t, tt.a, s)
			})
		})
	}
}
