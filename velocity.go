package unit

import (
	"math"

	"github.com/einride/can-databases/gen/go/dataspeed"
	"github.com/einride/can-databases/gen/go/emoss"
	"github.com/einride/plan"
	"github.com/einride/simulator"
)

// Speed is a value between 0 and 255 representing a vehicle speed (Km/h).
type Speed uint8

const SpeedUnit = "km/h"

// Speed constants
const (
	NoSpeed  Speed = 0
	MaxSpeed Speed = 255
)

const (
	wheelCircumference   = 2.131 // Meters
	metersPerMinuteToKmH = 60. / 1000.
	speedConstant        = 0.1 // Concluded from testing in vehicle
)

// TeleOperationSpeed is a float64 representing a unit-less speed quantity.
// Unit is provided to Teleoperator together with the quantity.
type TeleOperationSpeed = float64

// SpeedFromEMOSS returns the Speed as reported by the EMOSS CAN bus.
func SpeedFromEMOSS(speed emosscan.VMU1_VehicleSpeed) Speed {
	return Speed(speed)
}

// SpeedFromDataspeed returns the Speed as reported by the Dataspeed CAN bus.
func SpeedFromDataspeed(speed dataspeedcan.WheelSpeed_Report) Speed {
	avgSpeed := (int(speed.FR) + int(speed.FL) + int(speed.RL) + int(speed.RR)) / 4
	return Speed(int(float64(avgSpeed) * wheelCircumference * metersPerMinuteToKmH * speedConstant))
}

// ToTeleOperation returns the Speed converted for consumption by teleoperation GRPC interfaces.
func (speed Speed) ToTeleOperation() TeleOperationSpeed {
	return TeleOperationSpeed(speed)
}

func SimianVelocityToVehicle(vel *simulator.SimianVelocity) (speed Speed) {
	return Speed(math.Round(math.Sqrt(vel.Tx*vel.Tx+vel.Ty*vel.Ty+vel.Tz*vel.Tz) * 3.6))
}

func (speed Speed) ToPlanSpeed() plan.Speed {
	return plan.Speed(speed)
}

func SpeedFromPlan(vel plan.Speed) Speed {
	return Speed(vel)
}
