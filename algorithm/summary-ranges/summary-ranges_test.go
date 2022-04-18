package summary_ranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_summaryRanges(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		{
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := summaryRanges(tt.nums)
			assert.Equalf(t, tt.want, got, "summaryRanges() = %v, want %v", got, tt.want)
		})
	}
}
