package minimum_time_visiting_all_points

import (
	"testing"
)

func Test_minTimeToVisitAllPoints(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			points: [][]int{
				{1, 1},
				{3, 4},
				{-1, 0},
			},
			want: 7,
		},
		{
			points: [][]int{
				{3, 2},
				{-2, 2},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToVisitAllPoints(tt.points); got != tt.want {
				t.Errorf("minTimeToVisitAllPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
