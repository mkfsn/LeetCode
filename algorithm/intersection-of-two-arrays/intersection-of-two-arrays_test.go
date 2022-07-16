package intersection_of_two_arrays

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersection(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2},
		},
		{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersection(tt.nums1, tt.nums2)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
