package unit

import (
	"fmt"
	"testing"

	"github.com/einride/can-databases/gen/go/dataspeed"
	"github.com/einride/can-databases/gen/go/emoss"
	"github.com/einride/can-databases/gen/go/plc"
	"github.com/einride/proto/gen/go/voysys"
	"github.com/stretchr/testify/assert"
)

func TestGearFromEMOSS(t *testing.T) {
	for _, tt := range []struct {
		value    emosscan.VMU2
		expected Gear
	}{
		{value: emosscan.VMU2{VehicleDirection: emosscan.VMU2_VehicleDirection_Reverse}, expected: GearReverse},
		{value: emosscan.VMU2{VehicleDirection: emosscan.VMU2_VehicleDirection_Neutral}, expected: GearNeutral},
		{value: emosscan.VMU2{VehicleDirection: emosscan.VMU2_VehicleDirection_Drive}, expected: GearDrive},
		{value: emosscan.VMU2{VehicleDirection: emosscan.VMU2_VehicleDirection(2)}, expected: GearError},
		{
			value: emosscan.VMU2{
				VehicleDirection: emosscan.VMU2_VehicleDirection_Neutral,
				VehicleParkBrake: true,
			},
			expected: GearPark,
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, GearFromEMOSS(tt.value))
		})
	}
}

func TestGear_ToVoysys(t *testing.T) {
	for _, tt := range []struct {
		value    Gear
		expected voysyspb.Gear
	}{
		{value: GearNeutral, expected: voysyspb.Gear_GEAR_NEUTRAL},
		{value: GearReverse, expected: voysyspb.Gear_GEAR_REVERSE},
		{value: GearDrive, expected: voysyspb.Gear_GEAR_DRIVE},
		{value: GearPark, expected: voysyspb.Gear_GEAR_PARK},
		{value: GearNotAvailable, expected: voysyspb.Gear_GEAR_INVALID},
		{value: Gear(123), expected: voysyspb.Gear_GEAR_INVALID},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ToVoysys())
		})
	}
}

func TestGear_ToPLC(t *testing.T) {
	for _, tt := range []struct {
		value    Gear
		expected plccan.Longitudinal_Gear
	}{
		{value: GearNeutral, expected: plccan.Longitudinal_Gear_Neutral},
		{value: GearReverse, expected: plccan.Longitudinal_Gear_Reverse},
		{value: GearDrive, expected: plccan.Longitudinal_Gear_Drive},
		{value: GearPark, expected: plccan.Longitudinal_Gear_Park},
		{value: GearNotAvailable, expected: plccan.Longitudinal_Gear_NotAvailable},
		{value: GearError, expected: plccan.Longitudinal_Gear_Error},
		{value: Gear(123), expected: plccan.Longitudinal_Gear_Error},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ToPLC())
		})
	}
}

func TestGear_ParkBrakeToPLC(t *testing.T) {
	for _, tt := range []struct {
		value    Gear
		expected plccan.Longitudinal_ParkBrake
	}{
		// Engaged
		{value: GearPark, expected: plccan.Longitudinal_ParkBrake_Engaged},
		{value: GearNeutral, expected: plccan.Longitudinal_ParkBrake_Engaged},
		// Disengaged
		{value: GearReverse, expected: plccan.Longitudinal_ParkBrake_Disengaged},
		{value: GearDrive, expected: plccan.Longitudinal_ParkBrake_Disengaged},
		{value: GearNotAvailable, expected: plccan.Longitudinal_ParkBrake_Disengaged},
		{value: GearError, expected: plccan.Longitudinal_ParkBrake_Disengaged},
		{value: Gear(123), expected: plccan.Longitudinal_ParkBrake_Disengaged},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ParkBrakeToPLC())
		})
	}
}

func TestGearFromDataspeed(t *testing.T) {
	for _, tt := range []struct {
		value    dataspeedcan.Gear_Report_STATE
		expected Gear
	}{
		{value: dataspeedcan.Gear_Report_STATE_None, expected: GearNotAvailable},
		{value: dataspeedcan.Gear_Report_STATE_Park, expected: GearPark},
		{value: dataspeedcan.Gear_Report_STATE_Reverse, expected: GearReverse},
		{value: dataspeedcan.Gear_Report_STATE_Neutral, expected: GearNeutral},
		{value: dataspeedcan.Gear_Report_STATE_Drive, expected: GearDrive},
		{value: dataspeedcan.Gear_Report_STATE_Low, expected: GearDrive},
		{value: dataspeedcan.Gear_Report_STATE(123), expected: GearError},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, GearFromDataspeed(tt.value))
		})
	}
}

func TestGear_ToDataspeed(t *testing.T) {
	for _, tt := range []struct {
		value    Gear
		expected dataspeedcan.Gear_Report_STATE
	}{
		{value: GearError, expected: dataspeedcan.Gear_Report_STATE_None},
		{value: GearNotAvailable, expected: dataspeedcan.Gear_Report_STATE_None},
		{value: GearNeutral, expected: dataspeedcan.Gear_Report_STATE_Neutral},
		{value: GearReverse, expected: dataspeedcan.Gear_Report_STATE_Reverse},
		{value: GearPark, expected: dataspeedcan.Gear_Report_STATE_Park},
		{value: GearDrive, expected: dataspeedcan.Gear_Report_STATE_Drive},
		{value: Gear(123), expected: dataspeedcan.Gear_Report_STATE_None},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ToDataspeed())
		})
	}
}

func TestGear_ToDataspeedCAN(t *testing.T) {
	for _, tt := range []struct {
		value    Gear
		expected dataspeedcan.Gear_Command_GCMD
	}{
		{value: GearError, expected: dataspeedcan.Gear_Command_GCMD(dataspeedcan.Gear_Report_STATE_None)},
		{value: GearNotAvailable, expected: dataspeedcan.Gear_Command_GCMD(dataspeedcan.Gear_Report_STATE_None)},
		{value: GearNeutral, expected: dataspeedcan.Gear_Command_GCMD(dataspeedcan.Gear_Report_STATE_Neutral)},
		{value: GearReverse, expected: dataspeedcan.Gear_Command_GCMD(dataspeedcan.Gear_Report_STATE_Reverse)},
		{value: GearPark, expected: dataspeedcan.Gear_Command_GCMD(dataspeedcan.Gear_Report_STATE_Park)},
		{value: GearDrive, expected: dataspeedcan.Gear_Command_GCMD(dataspeedcan.Gear_Report_STATE_Drive)},
		{value: Gear(123), expected: dataspeedcan.Gear_Command_GCMD(dataspeedcan.Gear_Report_STATE_None)},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ToDataspeedCAN())
		})
	}
}

func TestGearFromVoysys(t *testing.T) {
	for _, tt := range []struct {
		value    voysyspb.Gear
		expected Gear
	}{
		{value: voysyspb.Gear_GEAR_NEUTRAL, expected: GearNeutral},
		{value: voysyspb.Gear_GEAR_REVERSE, expected: GearReverse},
		{value: voysyspb.Gear_GEAR_DRIVE, expected: GearDrive},
		{value: voysyspb.Gear_GEAR_PARK, expected: GearPark},
		{value: voysyspb.Gear_GEAR_INVALID, expected: GearError},
		{value: voysyspb.Gear(123), expected: GearError},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, GearFromVoysys(tt.value))
		})
	}
}
