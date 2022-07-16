package k_closest_points_to_origin

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{
			points: [][]int{
				{1, 3},
				{-2, 2},
			},
			k: 1,
			want: [][]int{
				{-2, 2},
			},
		},
		{
			points: [][]int{
				{3, 3},
				{5, -1},
				{-2, 4},
			},
			k: 2,
			want: [][]int{
				{-2, 4},
				{3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.points, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
