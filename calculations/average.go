package calculations

// calculate the average of the data in the file
func Average(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return float64(sum) / float64(len(data))
}
