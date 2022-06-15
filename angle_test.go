package unit

import (
	"math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestAngle_FromDegrees(t *testing.T) {
	assert.Equal(t, math.Pi*Radian, 180*Degree)
}

func TestAngle_ToDegrees(t *testing.T) {
	assert.Equal(t, 180.0, (Radian * math.Pi).Get(Degree))
}

func TestAngle_String(t *testing.T) {
	assert.Equal(t, "360°", (2 * math.Pi * Radian).String())
}

func TestAngle_Degrees(t *testing.T) {
	assert.Equal(t, (math.Pi * Radian).Degrees(), float64(180))
}

func TestAngle_FromRadians(t *testing.T) {
	assert.Equal(t, FromRadians(math.Pi), math.Pi*Radian)
}

func TestAngle_WrapMinusPiPi(t *testing.T) {
	type test struct {
		name        string
		angle, want Angle
	}
	tests := []test{
		{angle: 1.0, want: 1.0, name: "within positive"},
		{angle: -1.0, want: -1.0, name: "within negative"},
		{angle: -4.0, want: -4.0 + math.Pi*2, name: "negative change"},
		{angle: 4.0, want: 4.0 - math.Pi*2, name: "positive change"},
		{angle: -8.0, want: -8.0 + math.Pi*2, name: "double negative"},
		{angle: 8.0, want: 8.0 - math.Pi*2, name: "double positive"},
		{angle: -12.0, want: -12.0 + math.Pi*4, name: "triple negative"},
		{angle: 12.0, want: 12.0 - math.Pi*4, name: "triple positive"},
		{angle: math.Pi, want: -math.Pi, name: "pi"},
		{angle: -math.Pi, want: -math.Pi, name: "-pi"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := tc.angle.WrapMinusPiPi()
			assert.Assert(t, math.Abs(tc.want.Radians()-got.Radians()) < 1e-5)
		})
	}
}
