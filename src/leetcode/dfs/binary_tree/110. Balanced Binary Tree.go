package binary_tree

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils/mathhelper"
)

func isBalanced(root *mytreenode.TreeNode) bool {
	return isBalancedCore(root) != -1
}

func isBalancedCore(root *mytreenode.TreeNode) int {
	if root == nil {
		return 0
	}
	depthLeft := isBalancedCore(root.Left)
	depthRight := isBalancedCore(root.Right)
	if depthLeft == -1 || depthRight == -1 {
		return -1
	}
	if mathhelper.Abs(depthLeft-depthRight) > 1 {
		return -1
	}
	return mathhelper.Max(depthLeft, depthRight) + 1
}
