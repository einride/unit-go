package unit

import (
	"fmt"
	"testing"

	"github.com/einride/can-databases/gen/go/dataspeed"
	"github.com/einride/can-databases/gen/go/plc"
	"github.com/stretchr/testify/assert"
)

func TestThrottle_TeleOperationToPod(t *testing.T) {
	for _, tt := range []struct {
		value    TeleOperationThrottle
		expected Throttle
		error    bool
	}{
		{value: teleOperationNoThrottle, expected: NoThrottle},
		{value: teleOperationMaxThrottle, expected: MaxThrottle},
		{value: teleOperationMaxThrottle + 1, error: true},
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

func TestThrottle_PodToTeleOperation(t *testing.T) {
	for _, tt := range []struct {
		pod           Throttle
		teleOperation TeleOperationThrottle
	}{
		{pod: NoThrottle, teleOperation: teleOperationNoThrottle},
		{pod: MaxThrottle, teleOperation: teleOperationMaxThrottle},
		{pod: MaxThrottle + 1, teleOperation: teleOperationMaxThrottle},
	} {
		tt := tt
		t.Run(fmt.Sprintf("pod=%v,teleOperation=%v", tt.pod, tt.teleOperation), func(t *testing.T) {
			assert.Equal(t, tt.teleOperation, tt.pod.ToTeleOperation())
		})
	}
}
func TestThrottlePercentToPod(t *testing.T) {
	for _, tt := range []struct {
		value    float64
		expected Throttle
		error    bool
	}{
		{value: 100, expected: MaxThrottle},
		{value: 50, expected: MaxThrottle / 2},
		{value: 0, expected: NoThrottle},
		{value: 200, error: true},
		{value: -1, error: true},
	} {
		tt := tt
		t.Run(fmt.Sprintf(
			"Value=%v,Expected=%v,Error=%v", tt.value, tt.expected, tt.error), func(t *testing.T) {
			actual, err := ThrottlePercentToPod(tt.value)
			if tt.error {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestThrottle_ToPLC(t *testing.T) {
	for throttle := NoThrottle; throttle <= MaxThrottle; throttle++ {
		assert.Equal(t, plccan.Longitudinal_Throttle(throttle), throttle.ToPLC())
	}
}

func TestThrottle_ToDataspeed(t *testing.T) {
	for _, tt := range []struct {
		value    Throttle
		expected DataspeedThrottle
	}{
		{value: NoThrottle, expected: dataspeedNoThrottle},
		{value: MaxThrottle, expected: dataspeedMaxThrottle},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ToDataspeed())
		})
	}
}

func TestThrottle_ToDataspeedCAN(t *testing.T) {
	for _, tt := range []struct {
		value    Throttle
		expected dataspeedcan.Throttle_Command_PCMD
	}{
		{value: NoThrottle, expected: dataspeedcan.Throttle_Command_PCMD(dataspeedNoThrottle)},
		{value: MaxThrottle, expected: dataspeedcan.Throttle_Command_PCMD(dataspeedMaxThrottle)},
	} {
		tt := tt
		t.Run(fmt.Sprintf("value=%v,expected=%v", tt.value, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.ToDataspeedCAN())
		})
	}
}
