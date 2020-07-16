package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestDistance_Get(t *testing.T) {
	for _, tt := range []struct {
		msg      string
		d        Distance
		as       Distance
		expected float64
	}{
		{msg: "km as mile", d: Kilo * Metre, as: Mile, expected: 0.621371192},
		{msg: "mile as km", d: Mile, as: Kilo * Metre, expected: 1.609344000614692},
		{msg: "km as m", d: Kilo * Metre, as: Metre, expected: 1000},
	} {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.d.Get(tt.as))
		})
	}
}

func TestDistance_String(t *testing.T) {
	for _, tt := range []struct {
		d   Distance
		str string
	}{
		{d: Centi * Metre, str: "1cm"},
		{d: Kilo * Metre, str: "1km"},
		{d: 2.3 * Kilo * Metre, str: "2.3km"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				assert.Equal(t, tt.str, tt.d.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var d Distance
				assert.NilError(t, d.UnmarshalString(tt.str))
				assert.Equal(t, tt.d, d)
			})
		})
	}
}
