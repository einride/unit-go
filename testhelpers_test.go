package unit

import (
	"math"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestGetFloatingPointCmpOpts_ConvertToFromFloat32_Passes(t *testing.T) {
	// Take a large float64 value and add a diff that is too small to be
	// described in float32.
	float64Value := 1e30 + 100*math.Nextafter(1.0, 2.0)
	source := struct {
		Float64      float64
		Acceleration Acceleration
		Angle        Angle
		AngularSpeed AngularSpeed
		Curvature    Curvature
		Distance     Distance
		Force        Force
		Frequency    Frequency
		Jerk         Jerk
		Mass         Mass
		Speed        Speed
		Torque       Torque
	}{
		Float64:      float64Value,
		Acceleration: Acceleration(float64Value),
		Angle:        Angle(float64Value),
		AngularSpeed: AngularSpeed(float64Value),
		Curvature:    Curvature(float64Value),
		Distance:     Distance(float64Value),
		Force:        Force(float64Value),
		Frequency:    Frequency(float64Value),
		Jerk:         Jerk(float64Value),
		Mass:         Mass(float64Value),
		Speed:        Speed(float64Value),
		Torque:       Torque(float64Value),
	}
	truncated := source
	truncated.Float64 = float64(float32(source.Float64))
	truncated.Acceleration = Acceleration(float32(source.Acceleration))
	truncated.Angle = Angle(float32(source.Angle))
	truncated.AngularSpeed = AngularSpeed(float32(source.AngularSpeed))
	truncated.Curvature = Curvature(float32(source.Curvature))
	truncated.Distance = Distance(float32(source.Distance))
	truncated.Force = Force(float32(source.Force))
	truncated.Frequency = Frequency(float32(source.Frequency))
	truncated.Jerk = Jerk(float32(source.Jerk))
	truncated.Mass = Mass(float32(source.Mass))
	truncated.Speed = Speed(float32(source.Speed))
	truncated.Torque = Torque(float32(source.Torque))

	// Should not deepequal since float32 conversion loses precision.
	assert.Assert(t, !cmp.DeepEqual(source, truncated)().Success())
	assert.Check(t, source.Float64 != truncated.Float64)
	assert.Check(t, source.Acceleration != truncated.Acceleration)
	assert.Check(t, source.Angle != truncated.Angle)
	assert.Check(t, source.AngularSpeed != truncated.AngularSpeed)
	assert.Check(t, source.Curvature != truncated.Curvature)
	assert.Check(t, source.Distance != truncated.Distance)
	assert.Check(t, source.Force != truncated.Force)
	assert.Check(t, source.Frequency != truncated.Frequency)
	assert.Check(t, source.Jerk != truncated.Jerk)
	assert.Check(t, source.Mass != truncated.Mass)
	assert.Check(t, source.Speed != truncated.Speed)
	assert.Check(t, source.Torque != truncated.Torque)
	// Should deepequal with unit comparers option.
	assert.DeepEqual(t, source, truncated, GetFloatingPointCmpOpts()...)
}
