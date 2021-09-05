package parts

type Gear struct {
	Teeth int
	Angle float32
}

func (g *Gear) RotateByAngle(angle float32) {
	g.Angle += angle
	if g.Angle > 360{
		g.Angle -= 360
	}
}

func (g *Gear) RotateByTeeth(teeth int) {
	var rotation_percentage = float32(teeth / g.Teeth)
	g.Angle += rotation_percentage * 360
}