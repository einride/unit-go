package unit

import (
	"fmt"
	"testing"

	"github.com/einride/can-databases/gen/go/dataspeed"
	"github.com/einride/can-databases/gen/go/plc"
	"github.com/stretchr/testify/assert"
)

func TestBrake_TeleOperationToPod(t *testing.T) {
	for _, tt := range []struct {
		value    TeleOperationBrake
		expected Brake
		error    bool
	}{
		{value: teleOperationNoBrake, expected: NoBrake},
		{value: teleOperationMaxBrake, expected: MaxBrake},
		{value: teleOperationMaxBrake + 1, error: true},
	} {
		tt := tt
		t.Run(fmt.Sprintf(
			"TeleOperation=%v,Expected=%v,Error=%v",
			tt.value, tt.expected, tt.error), func(t *testing.T) {
			actual, err := tt.value.ConvertToPod()
			if tt.error {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestBrake_PodToTeleOperation(t *testing.T) {
	for _, tt := range []struct {
		pod           Brake
		teleOperation TeleOperationBrake
	}{
		{pod: NoBrake, teleOperation: teleOperationNoBrake},
		{pod: MaxBrake, teleOperation: teleOperationMaxBrake},
		{pod: MaxBrake + 1, teleOperation: teleOperationMaxBrake},
	} {
		tt := tt
		t.Run(fmt.Sprintf("pod=%v,teleOperation=%v", tt.pod, tt.teleOperation), func(t *testing.T) {
			assert.Equal(t, tt.teleOperation, tt.pod.ToTeleOperation())
		})
	}
}

func TestBrakePercentPLC(t *testing.T) {
	assert.Equal(t, MaxBrake, BrakePercentPLC(100))
	assert.Equal(t, MaxBrake, BrakePercentPLC(99.995428507))
	assert.Equal(t, NoBrake, BrakePercentPLC(0))
}

func TestBrake_ToPLC(t *testing.T) {
	for brake := NoBrake; brake <= MaxBrake; brake++ {
		assert.Equal(t, plccan.Longitudinal_Brake(brake), brake.ToPLC())
	}
}

func TestBrake_ToDataspeed(t *testing.T) {
	for _, tt := range []struct {
		value    Brake
		expected DataspeedBrake
	}{
		{value: NoBrake, expected: DataspeedNoBrake},
		{value: MaxBrake, expected: DataspeedMaxBrake},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ToDataspeed())
		})
	}
}

func TestBrake_DataspeedToPod(t *testing.T) {
	for _, tt := range []struct {
		value    DataspeedBrake
		expected Brake
	}{
		{value: DataspeedNoBrake, expected: NoBrake},
		{value: DataspeedMaxBrake, expected: MaxBrake},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ConvertToPod())
		})
	}
}

func TestBrake_ToDataspeedCAN(t *testing.T) {
	for _, tt := range []struct {
		value    Brake
		expected dataspeedcan.Brake_Command_PCMD
	}{
		{value: NoBrake, expected: dataspeedcan.Brake_Command_PCMD(DataspeedNoBrake)},
		{value: MaxBrake, expected: dataspeedcan.Brake_Command_PCMD(DataspeedMaxBrake)},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ToDataspeedCAN())
		})
	}
}
