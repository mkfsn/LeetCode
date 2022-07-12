package get_maximum_in_generated_array

import (
	"testing"
)

func Test_getMaximumGenerated(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 7, want: 3},
		{n: 2, want: 1},
		{n: 3, want: 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := getMaximumGenerated(tt.n); got != tt.want {
				t.Errorf("getMaximumGenerated() = %v, want %v", got, tt.want)
			}
		})
	}
}
