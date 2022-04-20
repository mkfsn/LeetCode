package kth_missing_positive_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKthPositive(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want int
	}{
		{
			arr:  []int{2, 3, 4, 7, 11},
			k:    5,
			want: 9,
		},
		{
			arr:  []int{1, 2, 3, 4},
			k:    2,
			want: 6,
		},
		{
			arr:  []int{9, 10},
			k:    1,
			want: 1,
		},
		{
			arr:  []int{1, 10},
			k:    1,
			want: 2,
		},
		{
			arr:  []int{1, 10},
			k:    8,
			want: 9,
		},
		{
			arr:  []int{1, 10},
			k:    19,
			want: 21,
		},
		{
			arr:  []int{100, 200},
			k:    19,
			want: 19,
		},
		{
			arr:  []int{100, 200},
			k:    100,
			want: 101,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := findKthPositive(tt.arr, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
