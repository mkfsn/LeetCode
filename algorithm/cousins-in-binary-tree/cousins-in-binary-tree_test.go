package cousins_in_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isCousins(t *testing.T) {
	type test struct {
		root *TreeNode
		x    int
		y    int
		want bool
	}

	tests := []func() test{
		func() test {
			nodes := []*TreeNode{
				{Val: 1},
				{Val: 2},
				{Val: 3},
				{Val: 4},
			}

			nodes[0].Left = nodes[1]
			nodes[0].Right = nodes[2]
			nodes[1].Left = nodes[3]

			return test{
				root: nodes[0],
				x:    4,
				y:    3,
				want: false,
			}
		},

		func() test {
			nodes := []*TreeNode{
				{Val: 1},
				{Val: 2},
				{Val: 3},
				{Val: 4},
				{Val: 5},
			}

			nodes[0].Left = nodes[1]
			nodes[0].Right = nodes[2]
			nodes[1].Right = nodes[3]
			nodes[2].Right = nodes[4]

			return test{
				root: nodes[0],
				x:    5,
				y:    4,
				want: true,
			}
		},

		func() test {
			nodes := []*TreeNode{
				{Val: 1},
				{Val: 2},
				{Val: 3},
				{Val: 4},
			}

			nodes[0].Left = nodes[1]
			nodes[0].Right = nodes[2]
			nodes[1].Right = nodes[3]

			return test{
				root: nodes[0],
				x:    2,
				y:    3,
				want: false,
			}
		},
	}

	for _, fn := range tests {
		tt := fn()
		t.Run("", func(t *testing.T) {
			got := isCousins(tt.root, tt.x, tt.y)
			assert.Equal(t, tt.want, got)
		})
	}
}
