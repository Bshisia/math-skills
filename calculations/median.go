package calculations

import "sort"

// calculate the median of the data provided
func Median(data []float64) float64 {
	sort.Float64s(data)
	len := len(data)
	if len%2 == 0 {
		return float64(data[(len-1)/2]+data[len/2]) / 2.0
	}
	return float64(data[len/2])
}
