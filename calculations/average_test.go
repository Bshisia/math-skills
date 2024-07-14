package calculations

import "testing"

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
			if got := Average(tt.args.data); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}
