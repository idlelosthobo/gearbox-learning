package main

import (
	"fmt"
	"learning/mover"
	"learning/parts"
)

func main() {
	newMotor := mover.Motor{1200, 5}
	newBox := parts.CreateRandomBox()
	newBox.RotateInputShaftByAngle(newMotor.GetAngleMovementBySeconds(60), newMotor.KW)
	fmt.Println(newBox.InputShaft.Rotations)
	fmt.Println(newBox.OutputShaft.Rotations)
}
