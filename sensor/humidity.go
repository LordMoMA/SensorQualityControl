package sensor

import (
	"fmt"
	"math"
)

type HumiditySensor struct {
	Name           string
	Readings       []float64
	ReferenceValue float64
}

func (h *HumiditySensor) ProcessReading(value float64) {
	h.Readings = append(h.Readings, value)
}

func (h *HumiditySensor) GetReadings() []float64 {
	return h.Readings
}

func (h *HumiditySensor) Evaluate() string {
	for _, reading := range h.Readings {
		if math.Abs(reading-h.ReferenceValue) > h.ReferenceValue*0.01 {
			return fmt.Sprintf("%s: discard", h.Name)
		}
	}
	return fmt.Sprintf("%s: OK", h.Name)
}

func (h *HumiditySensor) GetName() string {
	return h.Name
}
