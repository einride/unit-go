package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatWithPrefixAndSymbol(t *testing.T) {
	for _, tt := range []struct {
		value    float64
		symbol   string
		expected string
	}{
		{value: 1000000, symbol: "m", expected: "1000km"},
		{value: 2345, symbol: "m", expected: "2.345km"},
		{value: 1000, symbol: "m", expected: "1km"},
		{value: 999, symbol: "m", expected: "999m"},
		{value: 0.1, symbol: "m", expected: "0.1m"},
		{value: 0.09, symbol: "m", expected: "9cm"},
		{value: 0.01, symbol: "m/s", expected: "1cm/s"},
		{value: 0.009, symbol: "m", expected: "9mm"},
		{value: 0.001, symbol: "m", expected: "1mm"},
		{value: 0.0009, symbol: "m", expected: "900µm"},
		{value: 0.000001, symbol: "m", expected: "1µm"},
		{value: 0.000000999, symbol: "m", expected: "999nm"},
		{value: 0.000000001, symbol: "m", expected: "1nm"},
	} {
		tt := tt
		t.Run(tt.expected, func(t *testing.T) {
			require.Equal(t, tt.expected, formatWithPrefixAndSymbol(tt.value, tt.symbol))
		})
	}
}
