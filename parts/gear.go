package parts

type Gear struct {
	Teeth int
	Angle float32
}

func (g *Gear) RotateByAngle(angle float32) {
	g.Angle += angle
	if g.Angle > 360 {
		g.Angle -= 360
	}
}

func (g *Gear) RotateByTeeth(teeth int) {
	g.Angle += g.GetRotationByTeeth(teeth)
}

func (g *Gear) GetTeethByRotation(angle float32) int {
	anglePercentage := float32(angle / 360.0)
	return int(float32(g.Teeth) * anglePercentage)
}

func (g *Gear) GetRotationByTeeth(teeth int) float32 {
	rotationPercentage := float32(teeth) / float32(g.Teeth)
	angle := rotationPercentage * 360
	return angle
}
