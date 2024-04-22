package binary_tree

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils/mathhelper"
)

func maxDepth(root *mytreenode.TreeNode) int {
	if root == nil {
		return 0
	}
	return mathhelper.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
