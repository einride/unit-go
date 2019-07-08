package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatParse(t *testing.T) {
	for _, tt := range []struct {
		value  float64
		symbol string
		str    string
	}{
		{value: 1000000, symbol: "m", str: "1000km"},
		{value: 2345, symbol: "m", str: "2.345km"},
		{value: 1000, symbol: "m", str: "1km"},
		{value: 999, symbol: "m", str: "999m"},
		{value: 0.1, symbol: "m", str: "0.1m"},
		{value: 0.09, symbol: "m", str: "9cm"},
		{value: 0.01, symbol: "m/s", str: "1cm/s"},
		{value: 0.009, symbol: "m", str: "9mm"},
		{value: 0.001, symbol: "m", str: "1mm"},
		{value: 0.0009, symbol: "m", str: "900µm"},
		{value: 0.000001, symbol: "m", str: "1µm"},
		{value: 0.000000999, symbol: "m", str: "999nm"},
		{value: 0.000000001, symbol: "m", str: "1nm"},
		{value: -0.000000001, symbol: "m", str: "-1nm"},
		{value: -0.000000999, symbol: "m", str: "-999nm"},
		{value: -0.001, symbol: "m", str: "-1mm"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("format", func(t *testing.T) {
				require.Equal(t, tt.str, format(tt.value, tt.symbol))
			})
			t.Run("parse", func(t *testing.T) {
				parsed, err := parse(tt.str, map[string]float64{tt.symbol: 1})
				require.NoError(t, err)
				require.InDelta(t, tt.value, parsed, 1e-9)
			})
		})
	}
}

func TestFormat_Errors(t *testing.T) {
	for _, tt := range []struct {
		str   string
		units map[string]float64
		err   string
	}{
		{str: "", units: map[string]float64{"m": 1}, err: "parse '': unknown unit"},
		{str: "32rad", units: map[string]float64{"m": 1}, err: "parse '32rad': unknown unit"},
		{str: "m", units: map[string]float64{"m": 1}, err: "parse 'm': not a number"},
		{
			str:   "35qm",
			units: map[string]float64{"m": 1},
			err:   `parse '35qm': strconv.ParseFloat: parsing "35q": invalid syntax`,
		},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			parsed, err := parse(tt.str, tt.units)
			require.Equal(t, float64(0), parsed)
			require.NotNil(t, err)
			require.Equal(t, tt.err, err.Error())
		})
	}
}
