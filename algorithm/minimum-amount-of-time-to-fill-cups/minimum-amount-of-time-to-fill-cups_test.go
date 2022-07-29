package minimum_amount_of_time_to_fill_cups

import (
	"testing"
)

func Test_fillCups(t *testing.T) {
	tests := []struct {
		name   string
		amount []int
		want   int
	}{
		{
			amount: []int{1, 4, 2},
			want:   4,
		},
		{
			amount: []int{5, 4, 4},
			want:   7,
		},
		{
			amount: []int{5, 0, 0},
			want:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fillCups(tt.amount); got != tt.want {
				t.Errorf("fillCups() = %v, want %v", got, tt.want)
			}
		})
	}
}
