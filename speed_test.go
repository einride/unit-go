package unit

import (
	"math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSpeed_Kph(t *testing.T) {
	assert.Assert(t, math.Abs(3.6-MeterPerSecond(1).AsKilometerPerHour().KilometerPerHour()) < 1e-9)
}

func TestSpeed_UnitConversionMPSxKPH(t *testing.T) {
	assert.Equal(t, KilometerPerHour(1), KilometerPerHour(1).AsMetresPerSecond().AsKilometerPerHour())
}

func TestSpeed_Unmarshal(t *testing.T) {
	var m MeterPerSecond
	assert.NilError(t, m.UnmarshalString("36km/h"))
	assert.Equal(t, MeterPerSecond(10), m)
	var k KilometerPerHour
	assert.NilError(t, k.UnmarshalString("10m/s"))
	assert.Equal(t, KilometerPerHour(36), k)
}

func TestSpeed_String(t *testing.T) {
	for _, tt := range []struct {
		m   MeterPerSecond
		str string
	}{
		{m: 0, str: "0m/s"},
		{m: 1, str: "1m/s"},
		{m: 2.3 * Centi, str: "2.3cm/s"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.m.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var m MeterPerSecond
				assert.NilError(t, m.UnmarshalString(tt.str))
				assert.Equal(t, tt.m, m)
			})
		})
	}
}
