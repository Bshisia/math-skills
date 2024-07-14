package main

import (
	"os"
	"reflect"
	"testing"

	"statistics/calculations"
)

// Helper function to create temporary files for testing
func createTempFile(t *testing.T, content string) *os.File {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	return tmpfile
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    []float64
		wantErr bool
	}{
		{
			name:    "Valid data",
			content: "1.1\n2.2\n3.3",
			want:    []float64{1.1, 2.2, 3.3},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with the specified content
			tmpfile := createTempFile(t, tt.content)
			defer os.Remove(tmpfile.Name())

			// Call the ReadFile function with the temporary file's name
			got, err := calculations.ReadFile(tmpfile.Name())
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Average of positive integers",
			args: args{data: []float64{1, 2, 3, 4, 5}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculations.Average(tt.args.data); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "median of positive intergers",
			args: args{data: []float64{1, 2, 3, 4, 5}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculations.Median(tt.args.data); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariance(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Variance of positive intergers",
			args: args{data: []float64{1, 2, 3, 4, 5}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculations.Variance(tt.args.data); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Standard deviation of positive intergers",
			args: args{data: []float64{1, 2, 3, 4, 5}},
			want: 1.4142135623730951,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculations.StandardDeviation(tt.args.data); got != tt.want {
				t.Errorf("StandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
