package moveZeroes

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			nums: []int{0},
			want: []int{0},
		},
		{
			nums: []int{1},
			want: []int{1},
		},
		{
			nums: []int{1, 2, 3, 4, 0, 0},
			want: []int{1, 2, 3, 4, 0, 0},
		},

		{
			nums: []int{0, 0, 0, 0, 9, 8, 7},
			want: []int{9, 8, 7, 0, 0, 0, 0},
		},
		{
			nums: []int{0, 1},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)

			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
