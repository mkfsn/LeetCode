package last_stone_weight

import (
	"testing"
)

func Test_lastStoneWeight(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
		{
			stones: []int{1},
			want:   1,
		},
		{
			stones: []int{316, 157, 73, 106, 771, 828, 46, 212, 926, 604, 600, 992, 71, 51, 477, 869, 425, 405, 859, 924, 45, 187, 283, 590, 303, 66, 508, 982, 464, 398},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
