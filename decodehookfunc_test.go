package unit

import (
	"reflect"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDecodeHookFunc(t *testing.T) {
	for _, tt := range []struct {
		msg    string
		from   reflect.Type
		to     reflect.Type
		data   interface{}
		result interface{}
		err    string
	}{
		{
			msg:    "acceleration",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(Acceleration(0)),
			data:   "5m/s²",
			result: 5 * MetrePerSecondSquared,
		},
		{
			msg:  "acceleration error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(Acceleration(0)),
			data: "5",
			err:  "unmarshal acceleration: parse '5': unknown unit",
		},

		{
			msg:    "angle",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(Radians(0)),
			data:   "5°",
			result: Degree(5).AsRadians(),
		},
		{
			msg:  "angle error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(Degree(0)),
			data: "5",
			err:  "unmarshal angle: parse '5': unknown unit",
		},

		{
			msg:    "angular speed",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(AngularSpeed(0)),
			data:   "5rad/s",
			result: 5 * RadianPerSecond,
		},
		{
			msg:  "angular speed error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(AngularSpeed(0)),
			data: "5",
			err:  "unmarshal angular speed: parse '5': unknown unit",
		},

		{
			msg:    "distance",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(Distance(0)),
			data:   "5m",
			result: 5 * Metre,
		},
		{
			msg:  "distance error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(Distance(0)),
			data: "5",
			err:  "unmarshal distance: parse '5': unknown unit",
		},

		{
			msg:    "force",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(Force(0)),
			data:   "5N",
			result: 5 * Newton,
		},
		{
			msg:  "force error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(Force(0)),
			data: "5",
			err:  "unmarshal force: parse '5': unknown unit",
		},

		{
			msg:    "frequency",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(Frequency(0)),
			data:   "5Hz",
			result: 5 * Hertz,
		},
		{
			msg:  "frequency error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(Frequency(0)),
			data: "5",
			err:  "unmarshal frequency: parse '5': unknown unit",
		},

		{
			msg:    "jerk",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(Jerk(0)),
			data:   "5m/s³",
			result: 5 * MetrePerSecondCubed,
		},
		{
			msg:  "jerk error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(Jerk(0)),
			data: "5",
			err:  "unmarshal jerk: parse '5': unknown unit",
		},

		{
			msg:    "mass",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(Mass(0)),
			data:   "5kg",
			result: 5 * Kilo * Gram,
		},
		{
			msg:  "mass error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(Mass(0)),
			data: "5",
			err:  "unmarshal mass: parse '5': unknown unit",
		},

		{
			msg:    "speed",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(MeterPerSecond(0)),
			data:   "5m/s",
			result: MeterPerSecond(5),
		},
		{
			msg:  "speed error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(MeterPerSecond(0)),
			data: "5",
			err:  "unmarshal speed: parse '5': unknown unit",
		},

		{
			msg:    "torque",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(Torque(0)),
			data:   "5Nm",
			result: 5 * NewtonMetre,
		},
		{
			msg:  "torque error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(Torque(0)),
			data: "5",
			err:  "unmarshal torque: parse '5': unknown unit",
		},

		{
			msg:    "wrong to",
			from:   reflect.TypeOf(""),
			to:     reflect.TypeOf(""),
			data:   "5m/s²",
			result: "5m/s²",
		},
		{
			msg:    "wrong from",
			from:   reflect.TypeOf(5),
			to:     reflect.TypeOf(""),
			data:   5,
			result: 5,
		},
	} {
		actual, err := DecodeHookFunc(tt.from, tt.to, tt.data)
		if tt.err != "" {
			assert.Assert(t, is.Nil(actual))
			assert.Assert(t, err != nil)
			assert.Equal(t, tt.err, err.Error())
		} else {
			assert.NilError(t, err)
			assert.DeepEqual(t, tt.result, actual)
		}
	}
}
