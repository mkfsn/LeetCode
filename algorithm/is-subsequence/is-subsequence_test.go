package is_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
		{
			s:    "aaaaaa",
			t:    "bbaaaa",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsequence(tt.s, tt.t)
			assert.Equal(t, tt.want, got)
		})
	}
}
