package most_common_word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mostCommonWord(t *testing.T) {
	tests := []struct {
		paragraph string
		banned    []string
		want      string
	}{
		{
			paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
			banned:    []string{"hit"},
			want:      "ball",
		},
		{
			paragraph: "a.",
			banned:    []string{},
			want:      "a",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := mostCommonWord(tt.paragraph, tt.banned)
			assert.Equal(t, got, tt.want)
		})
	}
}
