package parts

import (
	"fmt"
	"learning/measurement"
)

type Box struct {
	InputShaft   Shaft
	OutputShaft  Shaft
	InternalTemp measurement.Temperature
	ExternalTemp measurement.Temperature
	Efficiency   float32
}

func CreateRandomBox() Box {
	randomBox := Box{
		Shaft{
			Gear{
				8,
				0.0,
			},
			0.0,
			0,
		},
		Shaft{
			Gear{
				32,
				0.0,
			},
			0.0,
			0,
		},
		measurement.Temperature{20},
		measurement.Temperature{20},
		0.97,
	}
	return randomBox
}

func (b *Box) GearRatio() int {
	return b.OutputShaft.Gear.Teeth / b.InputShaft.Gear.Teeth
}

func (b *Box) RotateInputShaftByAngle(angle float32, kwForce int, seconds float32) {
	b.InputShaft.RotateByAngle(angle)
	b.OutputShaft.RotateByAngle(b.OutputShaft.Gear.GetRotationByTeeth(b.InputShaft.Gear.GetTeethByRotation(angle)))
	b.ApplyForce(kwForce, seconds)
}

func (b *Box) ApplyForce(kwForce int, seconds float32) {
	forceTemperatureDelta := (1.0 - b.Efficiency) * float32(kwForce) * 10
	b.InternalTemp.ApplyTempDeltaByDegreesC(forceTemperatureDelta, seconds)
	b.InternalTemp.ApplyTempAffect(b.ExternalTemp, seconds)
}

func (b *Box) PrintStatus() {
	fmt.Printf("Input Rotations %d\n", b.InputShaft.Rotations)
	fmt.Printf("Output Rotations %d\n", b.OutputShaft.Rotations)
	fmt.Printf("Internal Temp %f\n", b.InternalTemp)
	fmt.Printf("External Temp %f\n", b.ExternalTemp)
}
