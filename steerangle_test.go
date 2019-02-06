package unit

import (
	"fmt"
	"testing"

	"github.com/einride/can-databases/gen/go/dataspeed"
	"github.com/einride/can-databases/gen/go/plc"
	"github.com/stretchr/testify/assert"
)

func TestSteerAngle_TeleOperation_ConvertToVehicle(t *testing.T) {
	for _, tt := range []struct {
		value    TeleOperationSteerAngle
		expected SteerAngle
		error    bool
	}{
		{value: teleOperationFullLeftSteerAngle, expected: FullLeftSteerAngle},
		{value: teleOperationFullRightSteerAngle, expected: FullRightSteerAngle},
		{value: teleOperationNeutralSteerAngle, expected: NeutralSteerAngle},
		{value: teleOperationFullRightSteerAngle + 1, error: true},
	} {
		tt := tt
		t.Run(fmt.Sprintf(
			"TeleOperation=%v,Expected=%v,Error=%v",
			tt.value, tt.expected, tt.error), func(t *testing.T) {
			actual, err := tt.value.ToVehicle()
			if tt.error {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestSteerAngle_PodToTeleOperation(t *testing.T) {
	for _, tt := range []struct {
		pod           SteerAngle
		teleOperation TeleOperationSteerAngle
	}{
		{
			pod:           FullLeftSteerAngle,
			teleOperation: teleOperationFullLeftSteerAngle,
		},
		{
			pod:           NeutralSteerAngle,
			teleOperation: teleOperationNeutralSteerAngle,
		},
		{
			pod:           FullRightSteerAngle,
			teleOperation: teleOperationFullRightSteerAngle,
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("pod=%v,teleOperation=%v", tt.pod, tt.teleOperation), func(t *testing.T) {
			assert.Equal(t, tt.teleOperation, tt.pod.ToTeleOperation())
		})
	}
}

func TestSteerAngle_SteerAngleDegreesVehicle(t *testing.T) {
	assert.Equal(t, FullRightSteerAngle, SteerAngleDegrees(55))
	assert.Equal(t, FullLeftSteerAngle, SteerAngleDegrees(-55))
	assert.Equal(t, NeutralSteerAngle, SteerAngleDegrees(0))
}

func TestSteerAngle_ToPLC(t *testing.T) {
	for steerAngle := FullLeftSteerAngle; steerAngle <= FullRightSteerAngle; steerAngle++ {
		assert.Equal(t, plccan.Latitudinal_SteerAngle(steerAngle), steerAngle.ToPLC())
	}
}

func TestSteerAngle_ToDataspeed(t *testing.T) {
	for _, tt := range []struct {
		value    SteerAngle
		expected DataspeedSteerAngle
	}{
		{value: FullLeftSteerAngle, expected: dataspeedFullLeftSteerAngle},
		{value: FullRightSteerAngle, expected: dataspeedFullRightSteerAngle},
		{value: NeutralSteerAngle, expected: dataspeedNeutralSteerAngle},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ToDataspeed())
		})
	}
}

func TestSteerAngle_AngleFromDataspeed(t *testing.T) {
	for _, tt := range []struct {
		value    dataspeedcan.Steering_Report_ANGLE
		expected SteerAngle
	}{
		{
			value:    dataspeedcan.Steering_Report_ANGLE(dataspeedFullLeftSteerAngle),
			expected: FullLeftSteerAngle,
		},
		{
			value:    dataspeedcan.Steering_Report_ANGLE(dataspeedFullRightSteerAngle),
			expected: FullRightSteerAngle,
		},
		{
			value:    dataspeedcan.Steering_Report_ANGLE(dataspeedNeutralSteerAngle),
			expected: NeutralSteerAngle,
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, AngleFromDataspeed(tt.value))
		})
	}
}

func TestSteerAngle_ToDataspeedCAN(t *testing.T) {
	for _, tt := range []struct {
		value    SteerAngle
		expected dataspeedcan.Steering_Command_SCMD
	}{
		{value: FullLeftSteerAngle, expected: dataspeedcan.Steering_Command_SCMD(dataspeedFullLeftSteerAngle)},
		{value: FullRightSteerAngle, expected: dataspeedcan.Steering_Command_SCMD(dataspeedFullRightSteerAngle)},
		{value: NeutralSteerAngle, expected: dataspeedcan.Steering_Command_SCMD(dataspeedNeutralSteerAngle)},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ToDataspeedCAN())
		})
	}
}
