package decrypt_string_from_alphabet_to_integer_mapping

import (
	"testing"
)

func Test_freqAlphabets(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			s:    "10#11#12",
			want: "jkab",
		},
		{
			s:    "1326#",
			want: "acz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := freqAlphabets(tt.s); got != tt.want {
				t.Errorf("freqAlphabets() = %v, want %v", got, tt.want)
			}
		})
	}
}
