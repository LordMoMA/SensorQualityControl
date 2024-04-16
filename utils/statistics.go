package utils

import (
	"math"
)

func Mean(values []float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func StdDev(values []float64, mean float64) float64 {
	var sqDiffSum float64
	for _, v := range values {
		sqDiffSum += math.Pow(v-mean, 2)
	}
	return math.Sqrt(sqDiffSum / float64(len(values)))
}

