package unit

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
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
			to:     reflect.TypeOf(Angle(0)),
			data:   "5°",
			result: 5 * Degree,
		},
		{
			msg:  "angle error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(Angle(0)),
			data: "5",
			err:  "unmarshal angle: parse '5': unknown unit",
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
			to:     reflect.TypeOf(Speed(0)),
			data:   "5m/s",
			result: 5 * MetrePerSecond,
		},
		{
			msg:  "speed error",
			from: reflect.TypeOf(""),
			to:   reflect.TypeOf(Speed(0)),
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
			require.Nil(t, actual)
			require.NotNil(t, err)
			require.Equal(t, tt.err, err.Error())
		} else {
			require.NoError(t, err)
			require.Equal(t, tt.result, actual)
		}
	}
}
