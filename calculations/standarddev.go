package calculations

import "math"

// calculate the standard deviation
func StandardDeviation(data []float64) float64 {
	variance := Variance(data)
	return math.Sqrt(variance)
}
