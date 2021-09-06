package measurement

const TemperatureDeltaStrength = 0.2

type Temperature struct {
	DegreesCelsius float32
}

func (t *Temperature) ApplyTempAffect(temperature Temperature, seconds float32) {
	temperatureDeltaNum := temperature.DegreesCelsius - t.DegreesCelsius
	t.DegreesCelsius = t.DegreesCelsius + temperatureDeltaNum*(TemperatureDeltaStrength*seconds)
}

func (t *Temperature) ApplyTempDeltaByDegreesC(degreesCelsius float32, seconds float32) {
	temperatureDelta := Temperature{t.DegreesCelsius + degreesCelsius}
	t.ApplyTempAffect(temperatureDelta, seconds)
}
