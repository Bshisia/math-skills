package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"statistics/calculations"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run [program name] [filename]")
		os.Exit(0)
	}
	Filename := os.Args[1]
	if !strings.HasSuffix(Filename, ".txt") {
		fmt.Println("Filename should be a txt file")
		os.Exit(1)
	}

	data, err := calculations.ReadFile(Filename)
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}
	if len(data) == 0 {
		fmt.Println("Empty data file")
		os.Exit(1)
	}

	average := calculations.Average(data)
	finalAverage := float64(math.Round(average))
	fmt.Printf("average: %.0f\n", finalAverage)
	Median := calculations.Median(data)
	finalMedian := float64(math.Round(Median))
	fmt.Printf("Median: %.0f\n", finalMedian)
	Variance := calculations.Variance(data)
	fialVariance := float64(math.Round(Variance))
	fmt.Printf("Variance: %.0f\n", fialVariance)
	standardDev := calculations.StandardDeviation(data)
	finalStandardDev := float64(math.Round(standardDev))
	fmt.Printf("Standard Deviation: %.0f\n", finalStandardDev)
}
