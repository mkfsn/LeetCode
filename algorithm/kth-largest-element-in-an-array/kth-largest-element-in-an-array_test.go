package kth_largest_element_in_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{nums: []int{3, 2, 1, 5, 6, 4}, k: 2},
			want: 5,
		},
		{
			args: args{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4},
			want: 4,
		},
		{
			args: args{nums: []int{6, 5, 4, 3, 2, 1}, k: 5},
			want: 2,
		},
		{
			args: args{nums: []int{7, 6, 5, 4, 3, 2, 1}, k: 5},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := findKthLargest(tt.args.nums, tt.args.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
