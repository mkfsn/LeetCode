package n_th_tribonacci_number

import (
	"testing"
)

func Test_tribonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		n    int
		want int
	}{
		// TODO: Add test cases.
		{n: 25, want: 1389537},
		{n: 4, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
