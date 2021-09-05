package main

import (
	"fmt"
	"learning/parts"
)

func main() {
	driveGear := parts.Gear{32, 0.0}
	driveGear.RotateByTeeth(16)
	fmt.Println(driveGear.Angle)
}