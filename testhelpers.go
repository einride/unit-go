package unit

import (
	"math"

	"github.com/google/go-cmp/cmp"
)

const floatComparisonTolerance float64 = 1e-7

// GetFloatingPointCmpOpts compares float unit types with floating point error
// tolerance. The function is useful when testing with conversions from unit
// types to/from float32 where precision is lost.
//
// Example usage:
// ```
// import "gotest.tools/v3/assert"
// [...]
// assert.DeepEqual(t, source, truncated, GetFloatingPointCmpOpts()...)
// ```.
func GetFloatingPointCmpOpts() []cmp.Option {
	compareFloat64 := func(x, y float64) bool {
		if x == 0 || y == 0 {
			return math.Abs(x) < floatComparisonTolerance && math.Abs(y) < floatComparisonTolerance
		}
		return math.Abs((x-y)/(x+y)) < floatComparisonTolerance
	}
	float64Comparer := cmp.Comparer(func(x, y float64) bool {
		return compareFloat64(x, y)
	})
	accelerationComparer := cmp.Comparer(func(x, y Acceleration) bool {
		return compareFloat64(x.MetersPerSecondSquared(), y.MetersPerSecondSquared())
	})
	angleComparer := cmp.Comparer(func(x, y Angle) bool {
		return compareFloat64(x.Radians(), y.Radians())
	})
	angularSpeedComparer := cmp.Comparer(func(x, y AngularSpeed) bool {
		return compareFloat64(x.RadiansPerSecond(), y.RadiansPerSecond())
	})
	curvatureComparer := cmp.Comparer(func(x, y Curvature) bool {
		return compareFloat64(x.PerMeters(), y.PerMeters())
	})
	distanceComparer := cmp.Comparer(func(x, y Distance) bool {
		return compareFloat64(x.Meters(), y.Meters())
	})
	forceComparer := cmp.Comparer(func(x, y Force) bool {
		return compareFloat64(x.Newtons(), y.Newtons())
	})
	frequencyComparer := cmp.Comparer(func(x, y Frequency) bool {
		return compareFloat64(x.Hertz(), y.Hertz())
	})
	jerkComparer := cmp.Comparer(func(x, y Jerk) bool {
		return compareFloat64(x.MetersPerSecondCubed(), y.MetersPerSecondCubed())
	})
	massComparer := cmp.Comparer(func(x, y Mass) bool {
		return compareFloat64(x.Kilograms(), y.Kilograms())
	})
	speedComparer := cmp.Comparer(func(x, y Speed) bool {
		return compareFloat64(x.MetersPerSecond(), y.MetersPerSecond())
	})
	torqueComparer := cmp.Comparer(func(x, y Torque) bool {
		return compareFloat64(x.NewtonMeters(), y.NewtonMeters())
	})

	return []cmp.Option{
		float64Comparer,
		accelerationComparer,
		angleComparer,
		angularSpeedComparer,
		curvatureComparer,
		distanceComparer,
		forceComparer,
		frequencyComparer,
		jerkComparer,
		massComparer,
		speedComparer,
		torqueComparer,
	}
}
