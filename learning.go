package main

import (
	"learning/mover"
	"learning/parts"
)

func main() {
	newMotor := mover.Motor{1200, 5}
	newBox := parts.CreateRandomBox()
	for i := float32(1.0); i <= 60.0; i++ {
		newBox.RotateInputShaftByAngle(newMotor.GetAngleMovementBySeconds(i), newMotor.KW, 1.0)
		newBox.PrintStatus()
	}
}
