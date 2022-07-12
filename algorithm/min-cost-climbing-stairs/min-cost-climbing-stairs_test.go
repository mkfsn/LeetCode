package min_cost_climbing_stairs

import (
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	tests := []struct {
		cost []int
		want int
	}{
		{
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minCostClimbingStairs(tt.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
