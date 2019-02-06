package unit

import (
	"github.com/einride/can-databases/gen/go/dataspeed"
	"github.com/einride/can-databases/gen/go/emoss"
	"github.com/einride/can-databases/gen/go/plc"
	"github.com/einride/proto/gen/go/voysys"
)

type Gear uint8

//go:generate stringer -type=Gear -trimprefix=Gear

const (
	GearPark Gear = iota
	GearNeutral
	GearReverse
	GearDrive
	GearNotAvailable
	GearError
)

func GearFromEMOSS(vmu2 emosscan.VMU2) Gear {
	if vmu2.VehicleParkBrake {
		return GearPark
	}
	switch vmu2.VehicleDirection {
	case emosscan.VMU2_VehicleDirection_Reverse:
		return GearReverse
	case emosscan.VMU2_VehicleDirection_Neutral:
		return GearNeutral
	case emosscan.VMU2_VehicleDirection_Drive:
		return GearDrive
	default:
		return GearError
	}
}

func (gear Gear) ToVoysys() voysyspb.Gear {
	switch gear {
	case GearNeutral:
		return voysyspb.Gear_GEAR_NEUTRAL
	case GearReverse:
		return voysyspb.Gear_GEAR_REVERSE
	case GearDrive:
		return voysyspb.Gear_GEAR_DRIVE
	case GearPark:
		return voysyspb.Gear_GEAR_PARK
	default:
		return voysyspb.Gear_GEAR_INVALID
	}
}

func (gear Gear) ToPLC() plccan.Longitudinal_Gear {
	switch gear {
	case GearNeutral:
		return plccan.Longitudinal_Gear_Neutral
	case GearReverse:
		return plccan.Longitudinal_Gear_Reverse
	case GearDrive:
		return plccan.Longitudinal_Gear_Drive
	case GearPark:
		return plccan.Longitudinal_Gear_Park
	case GearNotAvailable:
		return plccan.Longitudinal_Gear_NotAvailable
	case GearError:
		return plccan.Longitudinal_Gear_Error
	default:
		return plccan.Longitudinal_Gear_Error
	}
}

func (gear Gear) ParkBrakeToPLC() plccan.Longitudinal_ParkBrake {
	if gear == GearPark || gear == GearNeutral {
		return plccan.Longitudinal_ParkBrake_Engaged
	}
	return plccan.Longitudinal_ParkBrake_Disengaged
}

func GearFromDataspeed(gear dataspeedcan.Gear_Report_STATE) Gear {
	switch gear {
	case dataspeedcan.Gear_Report_STATE_None:
		return GearNotAvailable
	case dataspeedcan.Gear_Report_STATE_Park:
		return GearPark
	case dataspeedcan.Gear_Report_STATE_Reverse:
		return GearReverse
	case dataspeedcan.Gear_Report_STATE_Neutral:
		return GearNeutral
	case dataspeedcan.Gear_Report_STATE_Drive:
		return GearDrive
	case dataspeedcan.Gear_Report_STATE_Low: // TODO: implement crawl
		return GearDrive
	default:
		return GearError
	}
}

func (gear Gear) ToDataspeed() dataspeedcan.Gear_Report_STATE {
	switch gear {
	case GearError:
		return dataspeedcan.Gear_Report_STATE_None
	case GearNotAvailable:
		return dataspeedcan.Gear_Report_STATE_None
	case GearPark:
		return dataspeedcan.Gear_Report_STATE_Park
	case GearReverse:
		return dataspeedcan.Gear_Report_STATE_Reverse
	case GearNeutral:
		return dataspeedcan.Gear_Report_STATE_Neutral
	case GearDrive:
		return dataspeedcan.Gear_Report_STATE_Drive
	default:
		return dataspeedcan.Gear_Report_STATE_None
	}
}

func (gear Gear) ToDataspeedCAN() dataspeedcan.Gear_Command_GCMD {
	return dataspeedcan.Gear_Command_GCMD(gear.ToDataspeed())
}

func GearFromVoysys(gear voysyspb.Gear) Gear {
	switch gear {
	case voysyspb.Gear_GEAR_NEUTRAL:
		return GearNeutral
	case voysyspb.Gear_GEAR_REVERSE:
		return GearReverse
	case voysyspb.Gear_GEAR_DRIVE:
		return GearDrive
	case voysyspb.Gear_GEAR_PARK:
		return GearPark
	default:
		return GearError
	}
}
