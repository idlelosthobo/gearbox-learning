package main

import (
	"fmt"
	"learning/parts"
)

func main() {
	newBox := parts.CreateRandomBox()
	newBox.RotateInputShaftByAngle(250)
	fmt.Println(newBox.GetOutputShaftAngle())
}
