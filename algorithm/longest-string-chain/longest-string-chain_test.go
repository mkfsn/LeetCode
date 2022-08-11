package longest_string_chain

import (
	"testing"
)

func Test_longestStrChain(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			words: []string{
				"a", "b", "ba", "bca", "bda", "bdca",
			},
			want: 4,
		},
		{
			words: []string{
				"xbc", "pcxbcf", "xb", "cxbc", "pcxbc",
			},
			want: 5,
		},
		{
			words: []string{"abcd", "dbqca"},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestStrChain(tt.words); got != tt.want {
				t.Errorf("longestStrChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
