package binary_tree

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils"
)

func maxDepth(root *mytreenode.TreeNode) int {
	if root == nil {
		return 0
	}
	return utils.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
