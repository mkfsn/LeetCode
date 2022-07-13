package _sum_ii

import (
	"testing"
)

func Test_fourSumCount(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
		want  int
	}{
		{
			nums1: []int{1, 2},
			nums2: []int{-2, -1},
			nums3: []int{-1, 2},
			nums4: []int{0, 2},
			want:  2,
		},
		{
			nums1: []int{0},
			nums2: []int{0},
			nums3: []int{0},
			nums4: []int{0},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := fourSumCount(tt.nums1, tt.nums2, tt.nums3, tt.nums4); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
