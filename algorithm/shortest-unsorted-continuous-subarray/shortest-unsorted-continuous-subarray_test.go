package shortest_unsorted_continuous_subarray

import (
	"testing"
)

func Test_findUnsortedSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{2, 6, 4, 8, 10, 9, 15}},
			want: 5,
		},
		{
			args: args{nums: []int{1, 2, 3, 4}},
			want: 0,
		},
		{
			args: args{nums: []int{1}},
			want: 0,
		},
		{
			args: args{nums: []int{6, 4, 3, 2}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.args.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
