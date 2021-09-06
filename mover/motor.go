package mover

type Motor struct {
	RPM int
	KW  int
}

func (m *Motor) GetAngleMovementBySeconds(seconds float32) float32 {
	percentageOfMinute := seconds / 60
	rotations := float32(m.RPM) * percentageOfMinute
	return rotations * 360
}
