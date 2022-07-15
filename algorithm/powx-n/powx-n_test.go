package powx_n

import (
	"testing"
)

func Test_myPow(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		{
			x:    2,
			n:    10,
			want: 1024,
		},
		{
			x:    2.1,
			n:    3,
			want: 9.26100,
		},
		{
			x:    2,
			n:    -2,
			want: 0.25000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.x, tt.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
