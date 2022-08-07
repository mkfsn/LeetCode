package longest_repeating_character_replacement

import (
	"testing"
)

func Test_characterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			s:    "ABABABABBAAAAAAB",
			k:    3,
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.s, tt.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
