package parts

type Shaft struct {
	Gear  Gear
	Angle float32
}

func (s *Shaft) RotateByAngle(angle float32) {
	s.Angle += angle
	if s.Angle > 360 {
		s.Angle -= 360
	}
	s.Gear.RotateByAngle(angle)
}
