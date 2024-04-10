package sensor

import (
	"SensorQualityControl/utils"
	"fmt"
	"math"
)

type Thermometer struct {
	Name             string
	Readings         []float64
	KnownTemperature float64
}

func (t *Thermometer) ProcessReading(value float64) {
	t.Readings = append(t.Readings, value)
}

func (t *Thermometer) GetReadings() []float64 {
	return t.Readings
}

func (t *Thermometer) Evaluate() string {
	mean := utils.Mean(t.Readings)
	stdDev := utils.StdDev(t.Readings, mean)

	switch {
	case math.Abs(mean-t.KnownTemperature) <= 0.5 && stdDev < 3:
		return fmt.Sprintf("%s: ultra precise", t.Name)
	case math.Abs(mean-t.KnownTemperature) <= 0.5 && stdDev < 5:
		return fmt.Sprintf("%s: very precise", t.Name)
	default:
		return fmt.Sprintf("%s: precise", t.Name)
	}
}

func (t *Thermometer) GetName() string {
	return t.Name
}
