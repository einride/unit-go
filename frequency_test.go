package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				require.Equal(t, tt.str, tt.f.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s Frequency
				require.NoError(t, s.UnmarshalString(tt.str))
				require.Equal(t, tt.f, s)
			})
		})
	}
}
