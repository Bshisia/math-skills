package calculations

// calculate the variance of the provided data
func Variance(data []float64) float64 {
	Average := Average(data)
	var sum float64
	for _, value := range data {
		Difference := float64(value) - Average
		sum += Difference * Difference
	}
	return sum / float64(len(data))
}
