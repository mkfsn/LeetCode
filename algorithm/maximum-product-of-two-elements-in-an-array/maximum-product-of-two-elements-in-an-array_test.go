package maximum_product_of_two_elements_in_an_array

import (
	"testing"
)

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{3, 4, 5, 2},
			want: 12,
		},
		{
			nums: []int{1, 5, 4, 5},
			want: 16,
		},
		{
			nums: []int{3, 7},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
