package tree

import (
	"github.com/stretchr/testify/assert"
	"leetcode/src/define/treenode"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := &treenode.TreeNode{
		Val: 3,
		Left: &treenode.TreeNode{
			Val: 5,
			Left: &treenode.TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &treenode.TreeNode{
				Val: 2,
				Left: &treenode.TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &treenode.TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &treenode.TreeNode{
			Val: 1,
			Left: &treenode.TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &treenode.TreeNode{
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
