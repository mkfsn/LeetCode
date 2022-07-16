package container_with_most_water

import (
	"testing"
)

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			height: []int{1, 1},
			want:   1,
		},
		{
			height: []int{1, 1, 100, 100, 1, 1, 1, 1},
			want:   100,
		},
		{
			height: []int{1, 0, 0, 0, 0, 0, 0, 2, 2},
			want:   8,
		},
		{
			height: []int{1, 1, 2, 2, 0, 0, 0, 0, 0, 0, 0, 3, 3},
			want:   20,
		},
		{
			height: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:   25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
