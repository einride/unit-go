package unit

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestAngle_FromDegrees(t *testing.T) {
	assert.Equal(t, math.Pi*Radian, 180*Degree)
	assert.Equal(t, FromRadians(math.Pi), FromDegrees(180))
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

func TestAngle_WrapZeroTwoPi(t *testing.T) {
	const (
		twoPi = 2 * math.Pi
		epsi  = 1e-5
	)
	type test struct {
		name        string
		angle, want Angle
	}
	tests := []test{
		{angle: 1.0, want: 1.0, name: "within positive"},
		{angle: 4.0, want: 4.0, name: "within  positive (large)"},
		{angle: 8.0, want: 8.0 - twoPi, name: "double positive"},
		{angle: 12.0, want: 12.0 - twoPi, name: "double positive (large)"},
		{angle: 15.0, want: 15.0 - 2*twoPi, name: "triple positive"},
		{angle: -1.0, want: -1.0 + twoPi, name: "negative change"},
		{angle: -4.0, want: -4.0 + twoPi, name: "negative change (large)"},
		{angle: -8.0, want: -8.0 + 2*twoPi, name: "double negative"},
		{angle: -12.0, want: -12.0 + 2*twoPi, name: "double negative (large)"},
		{angle: -15.0, want: -15.0 + 3*twoPi, name: "triple negative"},
		{angle: math.Pi, want: math.Pi, name: "pi"},
		{angle: -math.Pi, want: math.Pi, name: "-pi"},
		{angle: twoPi, want: 0, name: "2 pi"},
		{angle: 0, want: 0, name: "0"},
	}
	var d float64
	withinEps := cmp.Comparer(func(x, y float64) bool { d = math.Abs(x - y); return d < epsi })
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := tc.angle.WrapZeroTwoPi()
			w, g := tc.want.Radians(), got.Radians()
			ok := cmp.Equal(w, g, withinEps)
			// requirement check
			assert.Assert(t, g < twoPi)
			assert.Assert(t, g >= 0)
			assert.Assert(t, math.Mod(g-tc.angle.Radians(), twoPi) < epsi)
			// exact check
			assert.Assert(t, ok, "got: %f, want: %f, diff: %f > %f", g, w, d, epsi)
		})
	}
}

func TestAngle_JSON(t *testing.T) {
	jsonString := "{\"TheImportantAngle\":0.786473}"
	type JSONStruct struct {
		TheImportantAngle Angle
	}
	var jsonStruct JSONStruct
	err := json.Unmarshal([]byte(jsonString), &jsonStruct)
	assert.NilError(t, err)
	assert.Equal(t, jsonStruct.TheImportantAngle.Radians(), 0.786473)

	marshaled, err := json.Marshal(jsonStruct)
	assert.NilError(t, err)
	assert.Equal(t, string(marshaled), jsonString)
}
