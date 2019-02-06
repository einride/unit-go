package unit

import (
	"math"
	"strconv"

	"github.com/einride/can-databases/gen/go/dataspeed"
	"github.com/einride/can-databases/gen/go/plc"
	"github.com/einride/plan"
	"github.com/einride/simulator"
	"github.com/pkg/errors"
)

// SteerAngle is a value between 9500 and 55000.
// To avoid unnecessary tearing of the steering servo the steering is limited between -55 and 55 degrees
//
//  9500     = -60 degrees
//  55000 =  60 degrees
//
//  11375 = -55 degrees
//  53083 =  55 degrees
type SteerAngle uint16

// TeleOperationSteerAngle is a value between 0 and 10000.
//
//  0     = full left
//  10000 = full right
type TeleOperationSteerAngle uint32

// DataspeedSteerAngle is a value between 4777 (max left) and -4670 (max right)
// To avoid unnecessary tearing of the steering servo, the steering is limited between -4500 and 4500
//
// -4500 = max right
// 4500 = max left
type DataspeedSteerAngle int16

// SteerAngle constants.
const (
	FullLeftSteerAngle  SteerAngle = 11374
	NeutralSteerAngle   SteerAngle = 32229
	FullRightSteerAngle SteerAngle = 53083
)

// TeleOperationSteerAngle constants.
const (
	teleOperationFullLeftSteerAngle  TeleOperationSteerAngle = 0
	teleOperationNeutralSteerAngle   TeleOperationSteerAngle = 5000
	teleOperationFullRightSteerAngle TeleOperationSteerAngle = 10000
)

const (
	dataspeedFullLeftSteerAngle  DataspeedSteerAngle = 4500
	dataspeedNeutralSteerAngle   DataspeedSteerAngle = 0
	dataspeedFullRightSteerAngle DataspeedSteerAngle = -4500
)

// ToVehicle converts a TeleOperationSteerAngle value to a vehicle SteerAngle value.
func (t TeleOperationSteerAngle) ToVehicle() (SteerAngle, error) {
	if t > teleOperationFullRightSteerAngle {
		return 0, errors.Errorf("steer angle %v out of bounds: [0, %v]", t, teleOperationFullRightSteerAngle)
	}
	return SteerAngle(math.Round(translate(
			float64(t),
			float64(teleOperationFullLeftSteerAngle),
			float64(teleOperationFullRightSteerAngle),
			float64(FullLeftSteerAngle),
			float64(FullRightSteerAngle)))),
		nil
}

// SteerAngleDegrees converts from a number between -60 and 60 to a SteerAngle value.
func SteerAngleDegrees(degrees float64) SteerAngle {
	var plcSteerAngle plccan.Latitudinal_SteerAngle
	plcSteerAngle.SetPhysical(float64(degrees))
	quantizedValue := SteerAngle(plcSteerAngle)
	if quantizedValue > FullRightSteerAngle {
		return FullRightSteerAngle
	}
	return quantizedValue
}

// SteerAngleFromPLC decodes a SteerAngle value received from the PLC.
func SteerAngleFromPLC(actualAngle plccan.Sensors_ActualAngle) SteerAngle {
	// SteerAngle uses the same quantized interval as the PLC
	return SteerAngle(actualAngle)
}

// ToPLC converts the SteerAngle value to a PLC steer angle value.
func (t SteerAngle) ToPLC() plccan.Latitudinal_SteerAngle {
	return plccan.Latitudinal_SteerAngle(t)
}

// Degrees converts SteerAngles to degrees.
func (t SteerAngle) Degrees() float64 {
	return t.ToPLC().Physical()
}

// ToTeleOperation converts the SteerAngle value to a TeleOperation steer angle value.
func (t SteerAngle) ToTeleOperation() TeleOperationSteerAngle {
	return TeleOperationSteerAngle(math.Round(translate(
		float64(t),
		float64(FullLeftSteerAngle), float64(FullRightSteerAngle),
		float64(teleOperationFullLeftSteerAngle), float64(teleOperationFullRightSteerAngle))))
}

// String returns a string representation of the SteerAngle.
func (t SteerAngle) String() string {
	return strconv.FormatUint(uint64(t), 10) +
		" (" + strconv.FormatFloat(t.ToPLC().Physical(), 'f', -1, 64) + "%)"
}

func (t SteerAngle) ToSimian() simulator.SimianSteerAngle {
	return simulator.SimianSteerAngle(translate(float64(t),
		float64(FullLeftSteerAngle), float64(FullRightSteerAngle),
		float64(simulator.SimianFullLeftSteerAngle), float64(simulator.SimianFullRightSteerAngle)))
}

func SimianSteerAngleToVehicle(angle simulator.SimianSteerAngle) (t SteerAngle) {
	return SteerAngle(math.Round(translate(float64(angle),
		float64(simulator.SimianFullLeftSteerAngle), float64(simulator.SimianFullRightSteerAngle),
		float64(FullLeftSteerAngle), float64(FullRightSteerAngle))))
}

func AngleFromDataspeed(t dataspeedcan.Steering_Report_ANGLE) SteerAngle {
	return SteerAngle(math.Round(translate(float64(t),
		float64(dataspeedFullRightSteerAngle), float64(dataspeedFullLeftSteerAngle),
		float64(FullRightSteerAngle), float64(FullLeftSteerAngle))))
}

func (t SteerAngle) ToDataspeed() DataspeedSteerAngle {
	return DataspeedSteerAngle(math.Round(translate(float64(t),
		float64(FullLeftSteerAngle), float64(FullRightSteerAngle),
		float64(dataspeedFullLeftSteerAngle), float64(dataspeedFullRightSteerAngle))))
}

func (t SteerAngle) ToDataspeedCAN() dataspeedcan.Steering_Command_SCMD {
	return dataspeedcan.Steering_Command_SCMD(t.ToDataspeed())
}

// SteerAngleFromPlan converts a plan.SteerAngle to a vehicle SteerAngle.
func SteerAngleFromPlan(angle plan.SteerAngleDegrees) SteerAngle {
	return SteerAngleDegrees(float64(angle))
}

func (t SteerAngle) SteerAngleToPlan() plan.SteerAngleDegrees {
	return plan.SteerAngleDegrees(t)
}
