package log

import (
	"SensorQualityControl/sensor"
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ParseLog(reader *bufio.Reader) []sensor.Sensor {
	sensors := make(map[string]sensor.Sensor)
	var knownTemperature, knownHumidity float64

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		parts := strings.Split(string(line), " ")
		if len(parts) < 2 {
			continue
		}

		switch parts[0] {
		case "reference":
			// Handle reference line
			knownTemperature, _ = strconv.ParseFloat(parts[1], 64)
			knownHumidity, _ = strconv.ParseFloat(parts[2], 64)
			continue
		case "thermometer":
			sensors[parts[1]] = &sensor.Thermometer{Name: parts[1], KnownTemperature: knownTemperature}
		case "humidity":
			sensors[parts[1]] = &sensor.HumiditySensor{Name: parts[1], ReferenceValue: knownHumidity}
		default:
			// Handle reading
			if sensor, ok := sensors[parts[1]]; ok {
				value, _ := strconv.ParseFloat(parts[2], 64)
				sensor.ProcessReading(value)
			}
		}
	}

	// Convert map to slice
	var sensorSlice []sensor.Sensor
	for _, s := range sensors {
		sensorSlice = append(sensorSlice, s)
	}

	return sensorSlice
}
