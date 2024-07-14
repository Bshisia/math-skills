package calculations

import (
	"fmt"
	"os"
	"math"
	"strconv"
	"strings"
)

func ReadFile(Filename string) ([]float64, error) {
	data, err := os.ReadFile(Filename)
	
	if err != nil {
		return nil, err
	}
	
	
	lines := strings.Split(string(data), "\n")
	// lines := strings.Trim(line, " ")
	
	var result []float64
	count := 0
	for _, line := range lines {
		count++
		line = strings.Trim(line, " ")
		if line == "" {
			continue
		}

		val, err := strconv.ParseFloat(line, 64)
		if val < math.MinInt64 || val > math.MaxInt64 {
			fmt.Printf("error: overflow value of %f incorporated at line %d\n", val, count)
			continue
		}
		if err != nil {
			fmt.Printf("error: non-integrer value %v incorporated at line %d\n", err, count)
			continue
		}
		result = append(result, val)
	}
	return result, nil
}
