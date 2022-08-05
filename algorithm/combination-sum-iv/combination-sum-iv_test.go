package combination_sum_iv

import (
	"testing"
)

func Test_combinationSum4(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{4, 2, 1},
			target: 32,
			want:   39882198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum4(tt.nums, tt.target); got != tt.want {
				t.Errorf("combinationSum4() = %v, want %v", got, tt.want)
			}
		})
	}
}
