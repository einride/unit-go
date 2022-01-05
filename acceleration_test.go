package unit

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestAcceleration_String(t *testing.T) {
	assert.Equal(t, "1.2345m/s²", (123.45 * Centi * MeterPerSecondSquared).String())
}
