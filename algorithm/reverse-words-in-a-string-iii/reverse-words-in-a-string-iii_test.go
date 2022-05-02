package reverse_words_in_a_string_iii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "Let's take LeetCode contest",
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			s:    "God Ding",
			want: "doG gniD",
		},
		{
			s:    "abc",
			want: "cba",
		},

		{
			s:    "abc    def       ghi",
			want: "cba    fed       ihg",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := reverseWords(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
