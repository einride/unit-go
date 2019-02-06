package unit

import (
	"testing"

	"github.com/einride/can-databases/gen/go/dataspeed"
	"github.com/stretchr/testify/assert"
)

func TestSpeedFromDataspeed(t *testing.T) {
	speed := dataspeedcan.WheelSpeed_Report{
		FL: 900,
		FR: 1100,
		RL: 1100,
		RR: 900,
	}
	s := SpeedFromDataspeed(speed)
	assert.Equal(t, 12, int(s))
}
