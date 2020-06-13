// Package common defines common components that will be used by both server and client.
package common

// MotorDriver is the driver name of the EV3 motor.
type MotorDriver string

const (
	LargeMotor  MotorDriver = "lego-ev3-l-motor"
	MediumMotor MotorDriver = "lego-ev3-m-motor"
)

// OutPort is the output port name.
type OutPort string

const (
	OutA OutPort = "ev3-ports:outA"
	OutB OutPort = "ev3-ports:outB"
	OutC OutPort = "ev3-ports:outC"
	OutD OutPort = "ev3-ports:outD"
)
