package binary_tree

import (
	"leetcode/src/define/mytreenode"
)

func invertTree(root *mytreenode.TreeNode) *mytreenode.TreeNode {
	if root == nil {
		return nil
	}
	right := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = right
	return root
}
