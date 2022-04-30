package sort_array_by_parity_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortArrayByParityII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{[]int{4, 2, 5, 7}},
			want: []int{2, 7, 4, 5},
		},
		{
			args: args{[]int{2, 3}},
			want: []int{2, 3},
		},
		{
			args: args{nums: []int{3, 1, 4, 2}},
			want: []int{2, 3, 4, 1},
		},
		{
			args: args{nums: []int{2, 3, 1, 1, 4, 0, 0, 4, 3, 3}},
			want: []int{0, 3, 2, 3, 4, 1, 4, 1, 0, 3},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := sortArrayByParityII(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
