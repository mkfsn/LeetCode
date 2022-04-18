package x_of_a_kind_in_a_deck_of_cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasGroupsSizeX(t *testing.T) {
	tests := []struct {
		deck []int
		want bool
	}{
		{
			deck: []int{1, 1},
			want: true,
		},
		{
			deck: []int{1},
			want: false,
		},
		{
			deck: []int{1, 2, 3, 4, 4, 3, 2, 1},
			want: true,
		},
		{
			deck: []int{1, 1, 1, 2, 2, 2, 3, 3},
			want: false,
		},
		{
			deck: []int{1, 1, 2, 2, 2, 2},
			want: true,
		},
		{
			deck: []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := hasGroupsSizeX(tt.deck)
			assert.Equal(t, got, tt.want)
		})
	}
}
