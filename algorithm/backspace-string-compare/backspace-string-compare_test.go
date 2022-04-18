package backspace_string_compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_backspaceCompare(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{
			s:    "ab#c",
			t:    "ad#c",
			want: true,
		},
		{
			s:    "ab##",
			t:    "c#d#",
			want: true,
		},
		{
			s:    "a#c",
			t:    "b",
			want: false,
		},
		{
			s:    "a##c",
			t:    "#a#c",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := backspaceCompare(tt.s, tt.t)
			assert.Equal(t, got, tt.want)
		})
	}
}
