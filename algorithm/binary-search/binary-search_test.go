package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			nums:   []int{5},
			target: 5,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := search(tt.nums, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}
