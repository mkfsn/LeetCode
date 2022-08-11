package find_greatest_common_divisor_of_array

import (
	"testing"
)

func Test_findGCD(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// TODO: Add test cases.
		{
			nums: []int{2, 5, 6, 9, 10},
			want: 2,
		},
		{
			nums: []int{3, 3},
			want: 3,
		},
		{
			nums: []int{7, 5, 6, 8, 3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGCD(tt.nums); got != tt.want {
				t.Errorf("findGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
