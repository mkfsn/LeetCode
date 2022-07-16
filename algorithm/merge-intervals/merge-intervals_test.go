package merge_intervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			intervals: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			intervals: [][]int{
				{1, 4},
				{4, 5},
			},
			want: [][]int{
				{1, 5},
			},
		},
		{
			intervals: [][]int{
				{1, 3},
				{4, 6},
				{3, 4},
			},
			want: [][]int{
				{1, 6},
			},
		},
		{
			intervals: [][]int{
				{1, 3},
				{3, 4},
				{6, 10},
				{11, 14},
				{1, 100},
			},
			want: [][]int{
				{1, 100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
