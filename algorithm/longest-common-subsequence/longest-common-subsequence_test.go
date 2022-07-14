package longest_common_subsequence

import (
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
		{
			text1: "ace",
			text2: "ace",
			want:  3,
		},
		{
			text1: "aaaa",
			text2: "aaba",
			want:  3,
		},
		{
			text1: "abc",
			text2: "def",
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
