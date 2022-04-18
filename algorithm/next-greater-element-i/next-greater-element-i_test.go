package next_greater_element_i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextGreaterElement(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},
		{
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := nextGreaterElement(tt.nums1, tt.nums2)
			assert.Equal(t, tt.want, got)
		})
	}
}
