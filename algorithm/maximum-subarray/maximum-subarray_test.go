package maximum_subarray

import (
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			nums: []int{0},
			want: 0,
		},
		{
			nums: []int{1},
			want: 1,
		},
		{
			nums: []int{-2, -2},
			want: -2,
		},
		{
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
