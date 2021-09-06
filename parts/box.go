package parts

type Box struct {
	InputShaft  Shaft
	OutputShaft Shaft
}

func CreateRandomBox() Box {
	randomBox := Box{
		Shaft{
			Gear{
				8,
				0.0,
			},
			0.0,
		},
		Shaft{
			Gear{
				32,
				0.0,
			},
			0.0,
		},
	}
	return randomBox
}

func (b *Box) GearRatio() int {
	return b.OutputShaft.Gear.Teeth / b.InputShaft.Gear.Teeth
}

func (b *Box) RotateInputShaftByAngle(angle float32) {
	b.InputShaft.RotateByAngle(angle)
	b.OutputShaft.Gear.RotateByTeeth(b.InputShaft.Gear.GetTeethByRotation(angle))
	b.OutputShaft.Angle = b.OutputShaft.Gear.Angle
}

func (b *Box) GetOutputShaftAngle() float32 {
	return b.OutputShaft.Angle
}
