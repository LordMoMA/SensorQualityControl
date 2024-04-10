package sensor

type Sensor interface {
	GetName() string
	ProcessReading(value float64)
	GetReadings() []float64
	Evaluate() string
}
