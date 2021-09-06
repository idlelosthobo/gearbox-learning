package parts

import "learning/measurement"

type Box struct {
	InputShaft   Shaft
	OutputShaft  Shaft
	InteralTemp  measurement.Temperature
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

func (b *Box) RotateInputShaftByAngle(angle float32, kwForce int) {
	b.InputShaft.RotateByAngle(angle)
	b.OutputShaft.RotateByAngle(b.OutputShaft.Gear.GetRotationByTeeth(b.InputShaft.Gear.GetTeethByRotation(angle)))
}

func (b *Box) ApplyForce(kwForce int) {
	forceTemperatureDelta := (1.0 - b.Efficiency) * float32(kwForce)
	b.InteralTemp.ApplyTempDeltaByDegreesC(forceTemperatureDelta, 1.0)
}
