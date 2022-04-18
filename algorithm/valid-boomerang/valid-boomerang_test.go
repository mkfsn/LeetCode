package valid_boomerang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isBoomerang(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		points [][]int
		want   bool
	}{
		{
			points: [][]int{
				{1, 1}, {2, 3}, {3, 2},
			},
			want: true,
		},
		{
			points: [][]int{
				{1, 1}, {2, 2}, {3, 3},
			},
			want: false,
		},
		{
			points: [][]int{
				{1, 1}, {2, 3}, {4, 3},
			},
			want: true,
		},
		{
			points: [][]int{
				{0, 0}, {0, 2}, {2, 1},
			},
			want: true,
		},
		{
			points: [][]int{
				{0, 0}, {2, 0}, {0, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := isBoomerang(tt.points)
			assert.Equal(t, tt.want, got)
		})
	}
}
