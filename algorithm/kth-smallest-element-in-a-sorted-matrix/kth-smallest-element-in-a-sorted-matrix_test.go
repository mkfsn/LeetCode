package kth_smallest_element_in_a_sorted_matrix

import (
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	tests := []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{
			matrix: [][]int{
				{1, 5, 9},
				{10, 11, 13},
				{12, 13, 15},
			},
			k:    8,
			want: 13,
		},
		{
			matrix: [][]int{
				{-5},
			},
			k:    1,
			want: -5,
		},
		{
			matrix: [][]int{
				{1, 2},
				{1, 3},
			},
			k:    2,
			want: 1,
		},
		{
			matrix: [][]int{
				{1, 1, 1, 1, 2},
				{1, 1, 1, 1, 2},
				{1, 1, 1, 1, 2},
				{1, 1, 1, 2, 3},
			},
			k:    10,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := kthSmallest(tt.matrix, tt.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
