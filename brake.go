package unit

import (
	"math"
	"strconv"

	"github.com/einride/can-databases/gen/go/dataspeed"
	"github.com/einride/can-databases/gen/go/plc"
	"github.com/einride/simulator"
	"github.com/pkg/errors"
)

// Brake is a value between 0 and 65533.
//
//  0     = 0%
//  65533 = 100%
type Brake uint16

// Brake constants.
const (
	NoBrake  Brake = 0
	MaxBrake Brake = 65533
)

// TeleOperationBrake is a value between 0 and 10000.
//
//  0     = 0%
//  10000 = 100%
type TeleOperationBrake uint32

// TeleOperation brake constants.
const (
	teleOperationNoBrake  TeleOperationBrake = 0
	teleOperationMaxBrake TeleOperationBrake = 10000
)

// DataspeedBrake is a value measured in the vehicle to be between 9000 and 20000.
//
// 9000 = 0%
// 20000 = 100%
type DataspeedBrake uint16

// Dataspeed brake constants.
const (
	DataspeedNoBrake        DataspeedBrake = 9000
	DataspeedMaxBrake       DataspeedBrake = 20000
	DataspeedBrakeForChange DataspeedBrake = 11000
)

// ConvertToPod converts a TeleOperationBrake to a Pod Brake value.
func (teleOperationBrake TeleOperationBrake) ConvertToPod() (Brake, error) {
	if teleOperationBrake > teleOperationMaxBrake {
		return 0, errors.Errorf(
			"brake %v out of bounds: [%v, %v]",
			teleOperationBrake, teleOperationNoBrake, teleOperationMaxBrake)
	}
	return Brake(math.Round(translate(
		float64(teleOperationBrake),
		float64(teleOperationNoBrake), float64(teleOperationMaxBrake),
		float64(NoBrake), float64(MaxBrake)))), nil
}

// BrakePercentPLC converts from a percentage (0%-100%) to a Brake value.
func BrakePercentPLC(percent float64) Brake {
	var plcBrake plccan.Longitudinal_Brake
	plcBrake.SetPhysical(percent)
	quantizedValue := Brake(plcBrake)
	// NOTE: 100% brake exceeds the maximum quantized interval
	if quantizedValue > MaxBrake {
		return MaxBrake
	}
	return quantizedValue
}

// ToPLC converts the Brake value to a PLC brake value.
func (brake Brake) ToPLC() plccan.Longitudinal_Brake {
	return plccan.Longitudinal_Brake(brake)
}

// ToTeleOperation converts the Brake value to a TeleOperationBrake value.
func (brake Brake) ToTeleOperation() TeleOperationBrake {
	return TeleOperationBrake(math.Round(translate(
		float64(brake),
		float64(NoBrake), float64(MaxBrake),
		float64(teleOperationNoBrake), float64(teleOperationMaxBrake))))
}

// String returns a string representation of the brake value.
func (brake Brake) String() string {
	return strconv.FormatUint(uint64(brake), 10) +
		" (" + strconv.FormatFloat(brake.ToPLC().Physical(), 'f', -1, 64) + "%)"
}

func (brake Brake) ToSimian() simulator.SimianBrake {
	return simulator.SimianBrake(translate(float64(brake),
		float64(NoBrake), float64(MaxBrake),
		float64(simulator.SimianNoBrake), float64(simulator.SimianMaxBrake)))
}

func SimianBrakeToVehicle(br simulator.SimianBrake) Brake {
	return Brake(math.Round(translate(float64(br),
		float64(simulator.SimianNoBrake), float64(simulator.SimianMaxBrake),
		float64(NoBrake), float64(MaxBrake))))
}

func (brake Brake) ToDataspeed() DataspeedBrake {
	return DataspeedBrake(math.Round(translate(float64(brake),
		float64(NoBrake), float64(MaxBrake),
		float64(DataspeedNoBrake), float64(DataspeedMaxBrake))))
}

func (brake DataspeedBrake) ConvertToPod() Brake {
	return Brake(math.Round(translate(float64(brake),
		float64(DataspeedNoBrake), float64(DataspeedMaxBrake),
		float64(NoBrake), float64(MaxBrake))))
}

func (brake Brake) ToDataspeedCAN() dataspeedcan.Brake_Command_PCMD {
	return dataspeedcan.Brake_Command_PCMD(brake.ToDataspeed())
}
