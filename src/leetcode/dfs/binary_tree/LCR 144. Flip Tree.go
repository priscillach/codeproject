package binary_tree

import (
	"leetcode/src/define/mytreenode"
)

func flipTree(root *mytreenode.TreeNode) *mytreenode.TreeNode {
	if root == nil {
		return nil
	}
	newNode := &mytreenode.TreeNode{
		Val:   root.Val,
		Left:  flipTree(root.Right),
		Right: flipTree(root.Left),
	}
	return newNode
}
