package power_of_four

import (
	"testing"
)

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{n: 1, want: true},
		{n: 0, want: false},
		{n: -4, want: false},
		{n: 5, want: false},
		{n: 6, want: false},
		{n: 7, want: false},
		{n: 8, want: false},
		{n: 16, want: true},
		{n: 1024, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.n); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
