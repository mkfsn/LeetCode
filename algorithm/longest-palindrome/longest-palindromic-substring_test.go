package longest_palindrome

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			s:    "babad",
			want: "bab",
		},
		{
			s:    "cbbd",
			want: "bb",
		},
		{
			s:    "abcddde",
			want: "ddd",
		},
		{
			s:    "abb",
			want: "bb",
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
