package can_make_arithmetic_progression_from_sequence

import (
	"testing"
)

func Test_canMakeArithmeticProgression(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			arr:  []int{3, 5, 1},
			want: true,
		},
		{
			arr:  []int{1, 2, 4},
			want: false,
		},
		{
			arr:  []int{4, 1, 4, 3},
			want: false,
		},
		{
			arr:  []int{0, 0, 0, 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakeArithmeticProgression(tt.arr); got != tt.want {
				t.Errorf("canMakeArithmeticProgression() = %v, want %v", got, tt.want)
			}
		})
	}
}
