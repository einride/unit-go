package unit

import (
	"math"
	"strconv"

	"github.com/einride/can-databases/gen/go/dataspeed"
	"github.com/einride/can-databases/gen/go/plc"
	"github.com/einride/simulator"
	"github.com/pkg/errors"
)

// Throttle is a value between 0 and 65533.
//
//  0     = 0%
//  65533 = 100%
type Throttle uint16

// Throttle constants.
const (
	NoThrottle  Throttle = 0
	MaxThrottle Throttle = 65533
)

// TeleOperationThrottle is a value between 0 and 10000.
//
//  0     = 0%
//  10000 = 100%
type TeleOperationThrottle uint32

// TeleOperation throttle constants.
const (
	teleOperationNoThrottle  TeleOperationThrottle = 0
	teleOperationMaxThrottle TeleOperationThrottle = 10000
)

// DataspeedThrottle is a value between 9827 and 52434
// Max throttle is tuned down for safety.
//
// 9827 = 0%
// 20000 = 100%
type DataspeedThrottle uint16

const (
	dataspeedNoThrottle  DataspeedThrottle = 9827
	dataspeedMaxThrottle DataspeedThrottle = 20000
)

// ConvertToPod converts a TeleOperation throttle value to a Pod throttle value.
func (teleOperationThrottle TeleOperationThrottle) ConvertToPod() (Throttle, error) {
	if teleOperationThrottle > teleOperationMaxThrottle {
		return 0, errors.Errorf(
			"throttle %v out of bounds: [%v, %v]",
			teleOperationThrottle, teleOperationNoThrottle, teleOperationMaxThrottle)
	}
	return Throttle(math.Round(translate(
		float64(teleOperationThrottle),
		float64(teleOperationNoThrottle), float64(teleOperationMaxThrottle),
		float64(NoThrottle), float64(MaxThrottle)))), nil
}

// ThrottlePercentToPod converts from a percentage (0%-100%) to a Throttle value.
func ThrottlePercentToPod(percent float64) (Throttle, error) {
	if percent > 100 {
		return MaxThrottle, errors.Errorf("Throttle out of bounds: [%v]", percent)
	}
	if percent < 0 {
		return NoThrottle, errors.Errorf("Throttle out of bounds: [%v]", percent)
	}
	return Throttle(percent/100*float64(MaxThrottle-NoThrottle) + float64(NoThrottle)), nil
}

// ToPLC converts the Throttle value to a PLC throttle value.
func (throttle Throttle) ToPLC() plccan.Longitudinal_Throttle {
	return plccan.Longitudinal_Throttle(throttle)
}

// ToTeleOperation converts the Throttle value to a TeleOperation throttle value.
func (throttle Throttle) ToTeleOperation() TeleOperationThrottle {
	return TeleOperationThrottle(math.Round(translate(
		float64(throttle),
		float64(NoThrottle), float64(MaxThrottle),
		float64(teleOperationNoThrottle), float64(teleOperationMaxThrottle))))
}

// String returns a string representation of the throttle value.
func (throttle Throttle) String() string {
	return strconv.FormatUint(uint64(throttle), 10) +
		" (" + strconv.FormatFloat(throttle.ToPLC().Physical(), 'f', -1, 64) + "%)"
}

func (throttle Throttle) ToSimian() simulator.SimianThrottle {
	return simulator.SimianThrottle(translate(float64(throttle),
		float64(NoThrottle), float64(MaxThrottle),
		float64(simulator.SimianNoThrottle), float64(simulator.SimianMaxThrottle)))
}

func SimianThrottleToVehicle(th simulator.SimianThrottle) (throttle Throttle) {
	return Throttle(math.Round(translate(float64(th),
		float64(simulator.SimianNoThrottle), float64(simulator.SimianMaxThrottle),
		float64(NoThrottle), float64(MaxThrottle))))
}

func (throttle Throttle) ToDataspeed() DataspeedThrottle {
	return DataspeedThrottle(math.Round(translate(float64(throttle),
		float64(NoThrottle), float64(MaxThrottle),
		float64(dataspeedNoThrottle), float64(dataspeedMaxThrottle))))
}

func (throttle Throttle) ToDataspeedCAN() dataspeedcan.Throttle_Command_PCMD {
	return dataspeedcan.Throttle_Command_PCMD(throttle.ToDataspeed())
}
