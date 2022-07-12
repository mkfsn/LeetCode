package permutations

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{3, 5, 7},
			want: [][]int{
				{3, 5, 7},
				{3, 7, 5},
				{5, 3, 7},
				{5, 7, 3},
				{7, 3, 5},
				{7, 5, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := permute(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
