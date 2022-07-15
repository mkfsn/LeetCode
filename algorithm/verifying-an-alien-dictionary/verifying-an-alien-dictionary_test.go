package verifying_an_alien_dictionary

import (
	"testing"
)

func Test_isAlienSorted(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		order string
		want  bool
	}{
		{
			words: []string{"hello", "leetcode"},
			order: "hlabcdefgijkmnopqrstuvwxyz",
			want:  true,
		},
		{
			words: []string{"word", "world", "row"},
			order: "worldabcefghijkmnpqstuvxyz",
			want:  false,
		},
		{
			words: []string{"apple", "app"},
			order: "abcdefghijklmnopqrstuvwxyz",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.words, tt.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
