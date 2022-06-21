package longest_increasing_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			nums: []int{0, 3, 1, 6, 2, 2, 7},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := lengthOfLIS(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
