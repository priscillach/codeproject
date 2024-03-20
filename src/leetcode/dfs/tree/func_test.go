package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := &mytreenode.TreeNode{
		Val: 3,
		Left: &mytreenode.TreeNode{
			Val: 5,
			Left: &mytreenode.TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &mytreenode.TreeNode{
				Val: 2,
				Left: &mytreenode.TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &mytreenode.TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &mytreenode.TreeNode{
			Val: 1,
			Left: &mytreenode.TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &mytreenode.TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	p := root.Left
	q := root.Right

	assert.Equal(t, lowestCommonAncestor(root, p, q).Val, 3)
}
