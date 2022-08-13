package special_array_with_x_elements_greater_than_or_equal_x

import (
	"testing"
)

func Test_specialArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{3, 5},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialArray(tt.nums); got != tt.want {
				t.Errorf("specialArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
