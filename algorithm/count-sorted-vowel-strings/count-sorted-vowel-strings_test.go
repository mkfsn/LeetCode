package count_sorted_vowel_strings

import (
	"testing"
)

func Test_countVowelStrings(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			n:    1,
			want: 5,
		},
		{
			n:    10,
			want: 1001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelStrings(tt.n); got != tt.want {
				t.Errorf("countVowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
