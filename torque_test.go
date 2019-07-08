package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTorque_String(t *testing.T) {
	for _, tt := range []struct {
		t   Torque
		str string
	}{
		{t: 0, str: "0Nm"},
		{t: 2.3 * Kilo * NewtonMetre, str: "2.3kNm"},
		{t: 3 * Milli * NewtonMetre, str: "3mNm"},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				require.Equal(t, tt.str, tt.t.String())
			})
			t.Run("unmarshal", func(t *testing.T) {
				var s Torque
				require.NoError(t, s.UnmarshalString(tt.str))
				require.Equal(t, tt.t, s)
			})
		})
	}
}
