package longest_palindromic_subsequence

import (
	"testing"
)

func Test_longestPalindromeSubseq(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			s:    "cbba",
			want: 2,
		},
		{
			s:    "cbbc",
			want: 4,
		},
		{
			s:    "abcdde",
			want: 2,
		},
		{
			s:    "bbbab",
			want: 4,
		},
		{
			s:    "abcdcdbdad", // abcdcba
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
