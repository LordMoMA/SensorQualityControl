package test

import (
	"SensorQualityControl/log"
	"SensorQualityControl/sensor"
	"bufio"
	"os"
	"testing"
)

func TestParseLog(t *testing.T) {
	file, err := os.Open("data/test1.txt")
	if err != nil {
		t.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	sensors := log.ParseLog(reader)

	// Check the number of sensors
	if len(sensors) != 4 {
		t.Errorf("Expected 4 sensors, got %d", len(sensors))
	}

	// Check the name of the first sensor
	if sensors[0].GetName() != "temp-1" {
		t.Errorf("Expected the first sensor to be 'temp-1', got '%s'", sensors[0].GetName())
	}

	// Check the first reading of the first sensor
	firstReading := sensors[0].GetReadings()[0]
	if firstReading != 72.4 {
		t.Errorf("Expected the first reading of the first sensor to be 72.4, got %f", firstReading)
	}

	// Check the known temperature of the first sensor
	thermometer, ok := sensors[0].(*sensor.Thermometer)
	if !ok {
		t.Errorf("Expected the first sensor to be a Thermometer")
	} else if thermometer.KnownTemperature != 70.0 {
		t.Errorf("Expected the known temperature of the first sensor to be 70.0, got %f", thermometer.KnownTemperature)
	}

	// Check the name of the second sensor
	if sensors[1].GetName() != "temp-2" {
		t.Errorf("Expected the second sensor to be 'temp-2', got '%s'", sensors[1].GetName())
	}

	// Check the first reading of the second sensor
	firstReading = sensors[1].GetReadings()[0]
	if firstReading != 69.5 {
		t.Errorf("Expected the first reading of the second sensor to be 69.5, got %f", firstReading)
	}

	// Check the known temperature of the second sensor
	thermometer, ok = sensors[1].(*sensor.Thermometer)
	if !ok {
		t.Errorf("Expected the second sensor to be a Thermometer")
	} else if thermometer.KnownTemperature != 70.0 {
		t.Errorf("Expected the known temperature of the second sensor to be 70.0, got %f", thermometer.KnownTemperature)
	}
}
