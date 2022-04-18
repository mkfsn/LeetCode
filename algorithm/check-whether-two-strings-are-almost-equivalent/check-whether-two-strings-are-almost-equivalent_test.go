package check_whether_two_strings_are_almost_equivalent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkAlmostEquivalent(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  bool
	}{
		{
			word1: "aaaa",
			word2: "bccb",
			want:  false,
		},
		{
			word1: "abcdeef",
			word2: "abaaacc",
			want:  true,
		},
		{
			word1: "cccddabba",
			word2: "babababab",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := checkAlmostEquivalent(tt.word1, tt.word2)
			assert.Equal(t, got, tt.want)
		})
	}
}
