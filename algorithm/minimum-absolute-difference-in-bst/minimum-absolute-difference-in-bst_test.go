package minimum_absolute_difference_in_bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMinimumDifference(t *testing.T) {
	type test struct {
		root *TreeNode
		want int
	}

	tests := []func() test{
		func() test {
			nodes := []*TreeNode{
				{Val: 1},
				{Val: 3},
				{Val: 2},
			}

			nodes[0].Right = nodes[1]
			nodes[1].Left = nodes[2]

			return test{
				root: nodes[0],
				want: 1,
			}
		},
		func() test {
			nodes := []*TreeNode{
				{Val: 236},
				{Val: 104},
				{Val: 701},
				{Val: 227},
				{Val: 911},
			}

			nodes[0].Left = nodes[1]
			nodes[0].Right = nodes[2]
			nodes[1].Right = nodes[3]
			nodes[2].Right = nodes[4]

			return test{
				root: nodes[0],
				want: 9,
			}
		},
	}
	for _, fn := range tests {
		t.Run("", func(t *testing.T) {
			tt := fn()
			got := getMinimumDifference(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}
