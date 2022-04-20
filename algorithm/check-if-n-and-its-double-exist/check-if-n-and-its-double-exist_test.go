package check_if_n_and_its_double_exist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkIfExist(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{
			arr:  []int{10, 2, 5, 3},
			want: true,
		},
		{
			arr:  []int{7, 1, 14, 11},
			want: true,
		},
		{
			arr:  []int{3, 1, 7, 11},
			want: false,
		},
		{
			arr:  []int{-2, 0, 10, -19, 4, 6, -8},
			want: false,
		},
		{
			arr:  []int{0, 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := checkIfExist(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
