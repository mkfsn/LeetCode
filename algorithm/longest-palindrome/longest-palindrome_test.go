package longest_palindrome

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			s:    "CCCCCdddcca",
			want: 9,
		},
		{
			s:    "a",
			want: 1,
		},
		{
			s:    "DD",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
