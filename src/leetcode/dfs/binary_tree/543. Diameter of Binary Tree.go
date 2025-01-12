package binary_tree

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils/mathhelper"
)

var diameterOfBinaryTreeMax int

// https://leetcode.com/problems/diameter-of-binary-tree/description/
// finish times: 2
func diameterOfBinaryTree(root *mytreenode.TreeNode) int {
	diameterOfBinaryTreeMax = 0
	diameterOfBinaryTreeCore(root)
	return diameterOfBinaryTreeMax
}

func diameterOfBinaryTreeCore(root *mytreenode.TreeNode) int {
	if root == nil {
		return 0
	}
	depthLeft := diameterOfBinaryTreeCore(root.Left)
	depthRight := diameterOfBinaryTreeCore(root.Right)
	diameterOfBinaryTreeMax = mathhelper.Max(diameterOfBinaryTreeMax, depthLeft+depthRight)
	return mathhelper.Max(depthLeft, depthRight) + 1
}
