package maximum_score_after_splitting_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxScore(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "011101",
			want: 5,
		},
		{
			s:    "00111",
			want: 5,
		},
		{
			s:    "1111",
			want: 3,
		},
		{
			s:    "0001",
			want: 4,
		},
		{
			s:    "0000",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := maxScore(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
