package find_smallest_letter_greater_than_target

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextGreatestLetter(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{
			letters: []byte{'c', 'f', 'j'},
			target:  'a',
			want:    'c',
		},
		{
			letters: []byte{'c', 'f', 'j'},
			target:  'c',
			want:    'f',
		},
		{
			letters: []byte{'c', 'f', 'j'},
			target:  'd',
			want:    'f',
		},
		{
			letters: []byte{'a', 'm'},
			target:  'z',
			want:    'a',
		},
		{
			letters: []byte{'a', 'm'},
			target:  'b',
			want:    'm',
		},
		{
			letters: []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'm'},
			target:  'b',
			want:    'm',
		},
		{
			letters: []byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'},
			target:  'e',
			want:    'n',
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := nextGreatestLetter(tt.letters, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}
