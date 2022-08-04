package cinema_seat_allocation

import (
	"testing"
)

func Test_maxNumberOfFamilies(t *testing.T) {
	tests := []struct {
		name          string
		n             int
		reservedSeats [][]int
		want          int
	}{
		{
			n: 3,
			reservedSeats: [][]int{
				{1, 2},
				{1, 3},
				{1, 8},
				{2, 6},
				{3, 1},
				{3, 10},
			},
			want: 4,
		},

		{
			n: 2,
			reservedSeats: [][]int{
				{2, 1},
				{1, 8},
				{2, 6},
			},
			want: 2,
		},

		{
			n: 4,
			reservedSeats: [][]int{
				{4, 3},
				{1, 4},
				{4, 6},
				{1, 7},
			},
			want: 4,
		},

		{
			n: 4,
			reservedSeats: [][]int{
				// [4,3],[1,4],[4,6],[1,7]
				{4, 3},
				{1, 4},
				{4, 6},
				{1, 7},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfFamilies(tt.n, tt.reservedSeats); got != tt.want {
				t.Errorf("maxNumberOfFamilies() = %v, want %v", got, tt.want)
			}
		})
	}
}
