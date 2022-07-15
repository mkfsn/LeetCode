package unique_morse_code_words

import (
	"testing"
)

func Test_uniqueMorseRepresentations(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			words: []string{"gin", "zen", "gig", "msg"},
			want:  2,
		},
		{
			words: []string{"a"},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueMorseRepresentations(tt.words); got != tt.want {
				t.Errorf("uniqueMorseRepresentations() = %v, want %v", got, tt.want)
			}
		})
	}
}
