package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
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
			require.Equal(t, tt.expected, tt.d.Get(tt.as))
		})
	}
}

func TestDistance_String(t *testing.T) {
	for _, tt := range []struct {
		d        Distance
		expected string
	}{
		{d: Centi * Metre, expected: "1cm"},
		{d: Kilo * Metre, expected: "1km"},
		{d: 2.3 * Kilo * Metre, expected: "2.3km"},
	} {
		tt := tt
		t.Run(tt.expected, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.d.String())
		})
	}
}
