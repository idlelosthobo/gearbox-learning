package parts

type Shaft struct {
	Gear      Gear
	Angle     float32
	Rotations int
}

func (s *Shaft) RotateByAngle(angle float32) {
	s.Angle += angle
	if s.Angle > 360 {
		s.Rotations += int(s.Angle) / 360
		s.Angle = float32(int(s.Angle) % 360)
	}
	s.Gear.RotateByAngle(angle)
}
