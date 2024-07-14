package calculations

import "testing"

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
			if got := StandardDeviation(tt.args.data); got != tt.want {
				t.Errorf("StandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
